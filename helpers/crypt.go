package helpers

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func Sha1(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", sha1.Sum(data))
}

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}
