package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	Version   string
	BuildTime string
)

func main() {
	action := flagParser()
	switch {
	case action == "version":
		displayVersion(Version, BuildTime)
		os.Exit(0)
	case action == "scrape":
		gvs.Scraper() //TODO: This doesn't go anywhere currently
	}
}
