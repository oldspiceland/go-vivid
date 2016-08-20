package mediatype

import ()

type Video struct {
	Media
	mediaFormat  string
	mediaRunTime string
}

func (v *Video) MediaFormat() string {
	return v.mediaFormat
}

func (v *Video) MediaRunTime() string {
	return v.mediaRunTime
}

func (v *Video) SetMediaFormat(newMediaFormat string) {
	v.mediaFormat = newMediaFormat
}

func (v *Video) SetMediaRunTime(newMediaRunTime string) {
	v.mediaRunTime = newMediaRunTime
}
