package main

import (
	"github.com/oldspiceland/go-vivid/gvutils"
	"os"
)

var (
	Version   string
	BuildTime string
)

func main() {
	action := gvutils.FlagParser()
	switch {
	case action == "version":
		gvutils.DisplayVersion(Version, BuildTime)
		os.Exit(0)
	case action == "scrape":
		//gvs.Scraper() //TODO: This doesn't go anywhere currently
	}
}
