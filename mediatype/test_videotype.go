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
