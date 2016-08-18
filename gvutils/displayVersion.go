package gvutils

import (
	"fmt"
)

var (
	var Version string
	var BuildTime string
)

func DisplayVersion(version, buildTime string) {
	fmt.Printf("Current Version is %s, built at %s.\n", Version, BuildTime)
}
