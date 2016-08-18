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
	//Begin Flag Setup
	verS := flag.Bool("v", false, "displays the current version and build time")
	verL := flag.Bool("version", false, "displays the current version and build time")
	//End Flag Setup & Parse
	flag.Parse()

	if *verS == true || *verL == true {
		displayVersion(Version, BuildTime)
		os.Exit(0)
	}

	fmt.Println("vim-go")
}
