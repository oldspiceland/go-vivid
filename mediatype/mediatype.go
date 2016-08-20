package mediatype

import ()

//This is a generic struct for holding things that are the same between different types
type Media struct {
	mediaName string
	mediaType string
	mediaSize int64
}

func MediaName() {
	return mediaName
}

func MediaType() {
	return mediaType
}

func MediaSize() {
	return mediaSize
}

func SetMediaName(newName string) {
	mediaName = newName
}

func SetMediaType(newType string) {
	mediaType = newType
}

func SetMediaSize(newSize int64) {
	mediaSize = newSize
}
