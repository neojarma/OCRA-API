package helper

import "golang.org/x/crypto/bcrypt"

func GetHashingPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
