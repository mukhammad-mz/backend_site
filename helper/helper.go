package helper

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GenerateToken() string {
	b := make([]byte, 50)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}

func GenerateMd5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// UUID Create new UUID
func UUID() string {
	return uuid.NewV4().String()
}

//GetDateTime return Day-Month-Year Hour:Minute:Second
func GetDateTime() string {
	return time.Now().Format("02-01-2006 15:04:05")
}

//GetDate return Day-Month-Year
func GetDate() string {
	return time.Now().Format("02-01-2006")
}

//GetTime return  Hour:Minute:Second
func GetTime() string {
	return time.Now().Format("15:04:05")
}
