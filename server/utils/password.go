package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func MakeHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {

	}

	return string(bytes)
}

func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)

	return err == nil
}

func IsHashed(str string) bool {
	return len(str) == 60
}
