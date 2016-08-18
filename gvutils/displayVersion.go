package gvutils

import (
	"fmt"
)

var (
	Version   string
	BuildTime string
)

func DisplayVersion() {
	fmt.Printf("Current Version is %s, built at %s.\n", Version, BuildTime)
}
