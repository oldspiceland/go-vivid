package gvutils

import (
	"flag"
)

//FlagParser contains the flag implentation for go-vivid, outputing a string "action"
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
