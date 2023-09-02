package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func Upload(bin []byte, path string) (url string, err error) {
	url, err = localUploadImpl(bin, path)
	fmt.Println(url, err)
	return
}

func IsExists(path string) bool {
	info, err := os.Stat(filepath.Join(
		"/home/afeather/Codes/golang/src/douyin/cmd/storage/static/",
		path))
	if err != nil {
		return false
	}
	return !info.IsDir()
}


func FileUrl(path string) string {
	return "http://192.168.1.110:8080/" + filepath.Join("static/", path)
}

func localUploadImpl(bin []byte, path string) (url string, err error) {
	err = os.WriteFile(filepath.Join(
		"/home/afeather/Codes/golang/src/douyin/cmd/storage/static/",
		path), bin, 0644)
	return FileUrl(path), err
}

