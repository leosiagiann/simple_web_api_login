package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	salt := 8
	pw := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(pw, salt)

	return string(hash)
}

func ComparePassword(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
