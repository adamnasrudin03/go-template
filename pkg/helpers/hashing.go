package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (res string, err error) {
	arrByte := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(arrByte, 10)
	return string(hash), err
}

func PasswordValid(h, p string) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
