package importer

import (
	"encoding/csv"
	"os"
	"strconv"
)

// TZ is a trimmed down representation of a timezone
type TZ struct {
	ID               int
	CountryCode      string
	Name             string
	Type             string
	UTCOffsetMins    int
	UTCOffsetMinsDST int
	AliasTo          string
	DeprecatedBy     string
}

// CSV reads tz.csv and generates SQL insert states
func CSV() ([]*TZ, error) {
	f, err := os.Open("./timezones/importer/tz.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	var tzs []*TZ
	for i, line := range lines {
		utcOffsetMins, err := strconv.Atoi(line[3])
		if err != nil {
			return nil, err
		}
		utcOffsetMinsDST, err := strconv.Atoi(line[4])
		if err != nil {
			return nil, err
		}

		tz := &TZ{
			ID:               i + 1,
			CountryCode:      line[0],
			Name:             line[1],
			Type:             line[2],
			UTCOffsetMins:    utcOffsetMins,
			UTCOffsetMinsDST: utcOffsetMinsDST,
		}

		if tz.Type == "Alias" {
			tz.AliasTo = line[5]
		} else if tz.Type == "Deprecated" {
			tz.DeprecatedBy = line[5]
		}

		tzs = append(tzs, tz)
	}
	return tzs, nil
}
