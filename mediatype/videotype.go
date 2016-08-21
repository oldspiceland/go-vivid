package mediatype

import ()

//Video is a struct that builds on the Media struct adding new attributes for Videos.
type Video struct {
	Media
	mediaFormat  string
	mediaRunTime string
}

//MediaFormat is a getter function that returns Video.mediaFormat.
func (v *Video) MediaFormat() string {
	return v.mediaFormat
}

//MediaRunTime is a getter function that returns Video.mediaRunTime.
func (v *Video) MediaRunTime() string {
	return v.mediaRunTime
}

//SetMediaFormat is a setter function that takes a string and assigns it to Video.mediaFormat
func (v *Video) SetMediaFormat(newMediaFormat string) {
	v.mediaFormat = newMediaFormat
}

//SetMediaRunTime is a setter function that takes a string and assigns it to Video.mediaRunTime
func (v *Video) SetMediaRunTime(newMediaRunTime string) {
	v.mediaRunTime = newMediaRunTime
}
