package importer

import (
	"encoding/csv"
	"os"
)

// TZ is a trimmed down representation of a timezone
type TZ struct {
	Name   string
	Format string
}

// CSV reads tz.csv and generates SQL insert states
func CSV() ([]TZ, error) {
	f, err := os.Open("./cmd/timezones/tz.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	var tzs []TZ
	for _, line := range lines {
		tzs = append(tzs, TZ{
			Name:   line[0],
			Format: line[1],
		})
	}
	return tzs, nil
}
