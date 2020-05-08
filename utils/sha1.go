package utils

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func ToSha1(s string) string {
	hash := sha1.New()
	_, err := io.WriteString(hash, s)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}
