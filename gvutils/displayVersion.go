package gvutils

import (
	"fmt"
)

var (
	//Version string assigned during build by linker passed flags
	Version string
	//BuildTime string assigned during build by linker passed flags
	BuildTime string
)

//DisplayVersion prints the version and buildtime strings to the console
func DisplayVersion() {
	fmt.Printf("Current Version is %s, built at %s.\n", Version, BuildTime)
}
