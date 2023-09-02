package utils

import (
	"io"
	"mime/multipart"
)


func ReadMultipart(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	bytes, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

