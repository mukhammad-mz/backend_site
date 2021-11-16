package helper

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// UUID Create new UUID
func UUID() string {
	return uuid.NewV4().String()
}

//GetDateTime return Day-Month-Year Hour:Minute:Second
func GetDateTime() string {
	return time.Now().Format("01-02-2006 15:04:05")
}

//GetDate return Day-Month-Year
func GetDate() string {
	return time.Now().Format("01-02-2006")
}

//GetTime return  Hour:Minute:Second
func GetTime() string {
	return time.Now().Format("15:04:05")
}
