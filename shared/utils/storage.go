package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"douyin/shared/initialize"

)

func Upload(bin []byte, hash string) (err error) {
	err = localUploadImpl(bin, hash)
	// 根据stack overflow的61283248号问题，使用%w更合适
	if err != nil {
		return fmt.Errorf("upload: %s %w", hash, err)
	}

	return nil
}

func IsExists(hash string) bool {
	info, err := os.Stat(filepath.Join(
		"../../cmd/storage/static/",
		hash))
	if err != nil {
		return false
	}
	return !info.IsDir()
}

var url = initialize.Config.GetString("video_srv_prefix")
func GetUrlFromHash(hash string) string {
	return url + filepath.Join("static/", hash)
}

func localUploadImpl(bin []byte, hash string) (err error) {
	err = os.WriteFile(filepath.Join(
		"../../cmd/storage/static/",
		hash), bin, 0644)
	return err
}
