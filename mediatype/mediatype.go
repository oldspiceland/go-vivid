package mediatype

import ()

//This is a generic struct for holding things that are the same between different types
type Media struct {
	mediaName string
	mediaType string
	mediaSize int64
}

func (m *Media) MediaName() string {
	return m.mediaName
}

func (m *Media) MediaType() string {
	return m.mediaType
}

func (m *Media) MediaSize() int64 {
	return m.mediaSize
}

func (m *Media) SetMediaName(newName string) {
	m.mediaName = newName
}

func (m *Media) SetMediaType(newType string) {
	m.mediaType = newType
}

func (m *Media) SetMediaSize(newSize int64) {
	m.mediaSize = newSize
}
