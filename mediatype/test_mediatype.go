package mediatype

import (
	"testing"
)

func setup() Media {
	m := new(Media)
	m.SetMediaName("Test Media")
	m.SetMediaType("Testable Media")
	m.SetMediaSize(1000000)
	return *m
}

func TestMediaName(t *testing.T) {
	n := setup()
	if n.MediaName() != "Test Media" {
		t.Errorf("expected Test Media got %s.\n", n.MediaName())
	}
}

func TestMediaType(t *testing.T) {
	n := setup()
	if n.MediaType() != "Testable Media" {
		t.Errorf("expected Testable Media got %s.\n", n.MediaType())
	}
}

func TestMediaSize(t *testing.T) {
	n := setup()
	if n.MediaSize != int64(1000000) {
		t.Errorf("expected 1000000 got %s.\n", n.MediaSize())
	}
}

func TestSetMediaName(t *testing.T) {
	n := setup()
	old := n.MediaName()
	n.SetMediaName("Set Test Media")
	if old != "Test Media" && n.MediaName() != "Set Test Media" {
		t.Error("Name not set correctly")
	}
}

func TestSetMediaType(t *testing.T) {
	n := setup()
	old := n.MediaType()
	n.SetMediaType("Set Testable Media")
	if old != "Testable Media" && n.MediaType() != "Set Testable Media" {
		t.Error("Type not set correctly")
	}
}

func TestSetMediaSize(t *testing.T) {
	n := setup()
	old := n.MediaSize()
	n.SetMediaSize(2000000)
	if old != 1000000 && n.MediaSize() != 2000000 {
		t.Error("Size not set correctly")
	}
}
