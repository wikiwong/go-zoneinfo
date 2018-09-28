package importer

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/wikiwong/go-zoneinfo/timezones/timex"
)

var zoneDir string

// OS reads the OS' timezone database and prints SQL Inserts
func OS() {
	// timezone db exists on diff location for diff os
	var zoneDirs = []string{
		"/usr/share/zoneinfo/",
		"/usr/share/lib/zoneinfo/",
		"/usr/lib/locale/TZ/",
	}
	for _, zoneDir = range zoneDirs {
		readFile("")
	}
}

func readFile(path string) {
	files, _ := ioutil.ReadDir(zoneDir + path)
	for _, f := range files {
		if f.Name() != strings.ToUpper(f.Name()[:1])+f.Name()[1:] {
			continue
		}

		if f.IsDir() {
			readFile(path + "/" + f.Name())
		} else {
			timezoneID := (path + "/" + f.Name())[1:]
			location, err := timex.LoadLocation(timezoneID)
			if err == nil {
				for _, zone := range location.Zones() {
					name, offset, isDST := zone.Extract()
					zoneTmpl := `%s - %s - %d - %v`
					fmt.Println(fmt.Sprintf(zoneTmpl, timezoneID, name, offset, isDST))
				}
			} else {
				fmt.Println("ERROR!", timezoneID, err)
			}
		}
	}
}
