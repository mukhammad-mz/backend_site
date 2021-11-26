package helper

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"time"
)

func GenerateMd5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenerateSha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func GetTimeID() string {
	return time.Now().Format("060102150405")
}
