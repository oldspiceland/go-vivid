package mediatype

import (
	"testing"
)

func setup_video() *Video {
	v := new(Video)
	v.mediaName = "Test Video"
	v.mediaType = "Testable Video"
	v.mediaSize = 1000000
	v.mediaFormat = "MPEG Test Video"
	v.mediaRunTime = "1 hour, 5 minutes"
	return v
}

func TestVideoFormat(t *testing.T) {
	n := setup_video()
	if n.MediaFormat() != "MPEG Test Video" {
		t.Errorf("Format incorrect, expected MPEG Test Video, got %s,\n", n.MediaFormat())
	}
}

func TestVideoRunTime(t *testing.T) {
	n := setup_video()
	if n.MediaRunTime() != "1 hour, 5 minutes" {
		t.Errorf("Time incorrect, expected 1 hour, 5 minutes, got %s,\n", n.MediaRunTime())
	}
}

func TestSetVideoFormat(t *testing.T) {
	n := setup_video()
	old := n.MediaFormat()
	n.SetMediaFormat("FLV Test Video")
	if old != "MPEG Test Video" && n.MediaFormat() != "FLV Test Video" {
		t.Error("Unable to set video format")
	}
}

func TestSetVideoRunTime(t *testing.T) {
	n := setup_video()
	old := n.MediaRunTime()
	n.SetMediaRunTime("2 hours, 10 minutes")
	if old != "1 hour, 5 minutes" && n.MediaRunTime() != "2 hours, 10 minutes" {
		t.Error("Unable to set run time")
	}
}

func TestVideoName(t *testing.T) {
	n := setup_video()
	if n.MediaName() != "Test Video" {
		t.Errorf("expected Test Media got %s.\n", n.MediaName())
	}
}

func TestVideoType(t *testing.T) {
	n := setup_video()
	if n.MediaType() != "Testable Video" {
		t.Errorf("expected Testable Media got %s.\n", n.MediaType())
	}
}

func TestVideoSize(t *testing.T) {
	n := setup_video()
	if n.MediaSize() != 1000000 {
		t.Errorf("expected 1000000 got %v.\n", n.MediaSize())
	}
}

func TestSetVideoName(t *testing.T) {
	n := setup_video()
	old := n.MediaName()
	n.SetMediaName("Set Test Video")
	if old != "Test Media" && n.MediaName() != "Set Test Video" {
		t.Error("Name not set correctly")
	}
}

func TestSetVideoType(t *testing.T) {
	n := setup_video()
	old := n.MediaType()
	n.SetMediaType("Set Testable Video")
	if old != "Testable Media" && n.MediaType() != "Set Testable Video" {
		t.Error("Type not set correctly")
	}
}

func TestSetVideoSize(t *testing.T) {
	n := setup_video()
	old := n.MediaSize()
	n.SetMediaSize(2000000)
	if old != 1000000 && n.MediaSize() != 2000000 {
		t.Error("Size not set correctly")
	}
}
