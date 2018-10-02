/**
 * Example usage formating "seed data" for a db
 */
package main

import (
	"fmt"

	"github.com/wikiwong/go-zoneinfo/timezones/importer"
)

func main() {
	tzs, err := importer.CSV()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, tz := range tzs {
		sqlTmpl := `(%d, '%s', '%s', %s, %d, %d, %s, %s, 'seed-data'),`
		fmt.Println(fmt.Sprintf(sqlTmpl, tz.ID, tz.Name, tz.Type, tz.CountryCode, tz.UTCOffsetMins, tz.UTCOffsetMinsDST, tz.AliasTo, tz.DeprecatedBy))
	}
}
