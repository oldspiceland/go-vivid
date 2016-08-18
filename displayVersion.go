package main

import (
	"fmt"
)

func displayVersion(version, buildTime string) {
	fmt.Printf("Current Version is %s, built at %s.\n", version, buildTime)
}
