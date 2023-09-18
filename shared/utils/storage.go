package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"douyin/shared/initialize"

)

func Upload(bin []byte, path string) (url string, err error) {
	url, err = localUploadImpl(bin, path)
	fmt.Println(url, err)
	return
}

func IsExists(path string) bool {
	info, err := os.Stat(filepath.Join(
		"../../cmd/storage/static/",
		path))
	if err != nil {
		return false
	}
	return !info.IsDir()
}

var url = initialize.Config.GetString("video_srv_prefix")
func FileUrl(path string) string {
	return url + filepath.Join("static/", path)
}

func localUploadImpl(bin []byte, path string) (url string, err error) {
	err = os.WriteFile(filepath.Join(
		"../../cmd/storage/static/",
		path), bin, 0644)
	return FileUrl(path), err
}

