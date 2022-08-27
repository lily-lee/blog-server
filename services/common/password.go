package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func EncodePassword(salt, password string) string {
	hash := hmac.New(sha256.New, []byte(salt))
	hash.Write([]byte(password))
	b := hash.Sum(nil)
	return fmt.Sprintf("%x", b)
}

func CheckPassword(salt, password, encodedStr string) bool {
	return EncodePassword(salt, password) == encodedStr
}
