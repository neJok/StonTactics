package imageutil

import (
	"mime/multipart"
	"net/http"
	"strings"
)

func IsValidImageFile(fileHeader *multipart.FileHeader) bool {
	// Open file and check magic number (first 512 bytes) to determine file type
	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	defer file.Close()

	buffer := make([]byte, 512) // Buffer size as per image file format requirements
	_, err = file.Read(buffer)
	if err != nil {
		return false
	}

	contentType := http.DetectContentType(buffer)
	return strings.HasPrefix(contentType, "image/")
}