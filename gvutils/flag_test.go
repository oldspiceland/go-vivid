package gvutils

import (
	"testing"
)

func TestVersionFlag(t *testing.T) {
	expectedStr := "Current Version is 0.0, built at 2016-08-20T12:30:00-0400"
	Version = "0.0"
	BuildTime = "2016-08-20T12:30:00-0400"
	result := DisplayVersion()
	if result != expectedStr {
		t.Fatalf("Expected result to be:\n%s, but got instead:\n%s\n", expectedStr, result)
	}
}
