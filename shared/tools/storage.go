package tools

import (
	"os"
	"path/filepath"
)
func Upload(bin []byte, path string) (url string, err error) {
	return localUploadImpl(bin, path)
}

func localUploadImpl(bin []byte, path string) (url string, err error) {
	err = os.WriteFile(filepath.Join(
		"/home/afeather/Codes/golang/src/douyin/cmd/storage/static/",
		path), bin, 0644)
	url = filepath.Join("http://localhost:8080/static/", path)
	return url, err
}
