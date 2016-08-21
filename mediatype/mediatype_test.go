package mediatype

import (
	"testing"
)

func setup() Media {
	m := new(Media)
	m.mediaName = "Test Media"
	m.mediaType = "Testable Media"
	m.mediaSize = 1000000
	return *m
}

//TestMediaName is a test function for Media.MediaName function.
func TestMediaName(t *testing.T) {
	n := setup()
	if n.MediaName() != "Test Media" {
		t.Errorf("expected Test Media got %s.\n", n.MediaName())
	}
}

//TestMediaType is a test function for Media.MediaType function.
func TestMediaType(t *testing.T) {
	n := setup()
	if n.MediaType() != "Testable Media" {
		t.Errorf("expected Testable Media got %s.\n", n.MediaType())
	}
}

//TestMediaSize is a test function for Media.MediaSize function.
func TestMediaSize(t *testing.T) {
	n := setup()
	if n.MediaSize() != int64(1000000) {
		t.Errorf("expected 1000000 got %s.\n", n.MediaSize())
	}
}

//TestSetMediaName is a test function for Media.SetMediaName function.
func TestSetMediaName(t *testing.T) {
	n := setup()
	old := n.MediaName()
	n.SetMediaName("Set Test Media")
	if old != "Test Media" && n.MediaName() != "Set Test Media" {
		t.Error("Name not set correctly")
	}
}

//TestSetMediaType is a test function for Media.SetMediaType function.
func TestSetMediaType(t *testing.T) {
	n := setup()
	old := n.MediaType()
	n.SetMediaType("Set Testable Media")
	if old != "Testable Media" && n.MediaType() != "Set Testable Media" {
		t.Error("Type not set correctly")
	}
}

//TestSetMediaSize is a test function for Media.SetMediaSize function.
func TestSetMediaSize(t *testing.T) {
	n := setup()
	old := n.MediaSize()
	n.SetMediaSize(2000000)
	if old != 1000000 && n.MediaSize() != 2000000 {
		t.Error("Size not set correctly")
	}
}
