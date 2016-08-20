package mediatype

import (
	"testing"
)

func init() Video {
	v := new(Video)
	v.mediaFormat = "MPEG Test Video"
	v.mediaRunTime = "1 hour, 5 minutes"
	return v
}

func TestMediaFormat(t *testing.T) {
	n := init()
	if n.MediaFormat() != "MPEG Test Video" {
		t.Errorf("Format incorrect, expected MPEG Test Video, got %s,\n", n.MediaFormat())
	}
}

func TestMediaRunTime(t *testing.T) {
	n := init()
	if n.MediaRunTime() != "1 hour, 5 minutes" {
		t.Errorf("Time incorrect, expected 1 hour, 5 minutes, got %s,\n", n.MediaRunTime())
	}
}

func TestSetMediaFormat(t *testing.T) {
	n := init()
	old := n.MediaFormat()
	n.SetMediaFormat("FLV Test Video")
	if old != "MPEG Test Video" && n.MediaFormat() != "FLV Test Video" {
		t.Error("Unable to set video format")
	}
}

func TestSetMediaRunTime(t *testing.T) {
	n := init()
	old := n.MediaRunTime()
	n.SetMediaRunTime("2 hours, 10 minutes")
	if old != "1 hour, 5 minutes" && n.MediaRunTime() != "2 hours, 10 minutes" {
		t.Error("Unable to set run time")
	}
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
	n := init()
	if n.MediaSize != 1000000 {
		t.Errorf("expected 1000000 got %s.\n", n.MediaSize())
	}
}

func TestSetMediaName(t *testing.T) {
	n := init()
	old := n.MediaName()
	n.SetMediaName("Set Test Media")
	if old != "Test Media" && n.MediaName() != "Set Test Media" {
		t.Error("Name not set correctly")
	}
}

func TestSetMediaType(t *testing.T) {
	n := init()
	old := n.MediaType()
	n.SetMediaType("Set Testable Media")
	if old != "Testable Media" && n.MediaType() != "Set Testable Media" {
		t.Error("Type not set correctly")
	}
}

func TestSetMediaSize(t *testing.T) {
	n := init()
	old := n.MediaSize()
	n.SetMediaSize(2000000)
	if old != 1000000 && n.MediaSize() != 2000000 {
		t.Error("Size not set correctly")
	}
}
