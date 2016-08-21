package mediatype

import ()

//Media is a generic struct for holding things that are the same between different types
type Media struct {
	mediaName string
	mediaType string
	mediaSize int64
}

//MediaName is a getter method which returns Media.mediaName string
func (m *Media) MediaName() string {
	return m.mediaName
}

//MediaType is a getter method which returns Media.mediaType string
func (m *Media) MediaType() string {
	return m.mediaType
}

//MediaSize is a getter method which returns Media.mediaSize int64
func (m *Media) MediaSize() int64 {
	return m.mediaSize
}

//SetMediaName is a setter method which takes a string and assigns it to Media.mediaName
func (m *Media) SetMediaName(newName string) {
	m.mediaName = newName
}

//SetMediaType is a setter method which takes a string and assigns it to Media.mediaType
func (m *Media) SetMediaType(newType string) {
	m.mediaType = newType
}

//SetMediaSize is a setter method which takes an int64 and assigns it to Media.mediaSize
func (m *Media) SetMediaSize(newSize int64) {
	m.mediaSize = newSize
}
