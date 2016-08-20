package gvutils

import ()

func ExampleVersionFlag() {
	Version = "0.0"
	BuildTime = "2016-08-20T12:30:00-0400"
	DisplayVersion()
	// Output: Current Version is 0.0, built at 2016-08-20T12:30:00-0400.
}
