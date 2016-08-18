package gvutils

import (
	"flag"
)

func FlagParser() (action string) {
	verS := flag.Bool("v", false, "displays the current version and build time")
	verL := flag.Bool("version", false, "displays the current version and build time")
	flag.Parse()

	if *verS || *verL == true {
		action = "version"
		return
	}
	return
}
