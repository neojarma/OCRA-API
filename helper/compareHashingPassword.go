package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashingPassword(hashedPassword, plainPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}
