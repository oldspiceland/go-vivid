package gvutils

import (
	"fmt"
)

func DisplayVersion(version, buildTime string) {
	fmt.Printf("Current Version is %s, built at %s.\n", version, buildTime)
}
