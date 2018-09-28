package main

import (
	"fmt"

	"github.com/wikiwong/go-zoneinfo/timezones/importer"
)

func main() {
	importer.OS()
}

// this has nothing to do with zoneinfo, it reads a csv file of {timezone},{format}
// (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
func getFromCSV() {
	tzs, err := importer.CSV()
	if err != nil {
		fmt.Println(err)
	}

	zoneTmpl := `%s - %s`
	for _, tz := range tzs {
		fmt.Println(fmt.Sprintf(zoneTmpl, tz.Name, tz.Format))
	}
}
