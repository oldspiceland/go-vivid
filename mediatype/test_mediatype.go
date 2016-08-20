package mediatype

import (
	"testing"
)

func init() Media {
	m := new(Media)
	m.SetMediaName("Test Media")
	m.SetMediaType("Testable Media")
	m.SetMediaSize(1000000)
	return m
}

func TestMediaName(t *testing.T) {
	n := init()
	if n.MediaName() != "Test Media" {
		t.Errorf("expected Test Media got %s.\n", n.MediaName())
	}
}

func TestMediaType(t *testing.T) {
	n := init()
	if n.MediaType() != "Testable Media" {
		t.Errorf("expected Testable Media got %s.\n", n.MediaType())
	}
}

func TestMediaSize(t *testing.T) {
	N := init()
	if n.MediaSize != 1000000 {
		t.Errorf("expected 1000000 got %s.\n", n.MediaSize())
	}
}
