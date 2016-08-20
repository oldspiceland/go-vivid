package mediatype

import ()

type Video struct {
	Media
	mediaFormat  string
	mediaRunTime string
}

func MediaFormat() {
	return mediaFormat
}

func MediaRunTime() {
	return mediaRunTime
}

func SetMediaFormat(newMediaFormat string) {
	mediaFormat = newMediaFormat
}

func SetMediaRunTime(newMediaRunTime string) {
	mediaRunTime = newMediaRunTime
}
