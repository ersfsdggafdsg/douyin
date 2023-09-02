package utils

import (
	"fmt"

	sha256 "github.com/minio/sha256-simd"
)

func SHA256(bytes []byte) string {
	hash := sha256.New()
	hash.Write(bytes)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
