package users

import (
	"crypto/rand"
	"fmt"
)

func NewToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
