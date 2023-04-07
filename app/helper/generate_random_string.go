package helper

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

func GetRandomString(length int) string {
	strings := []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890")
	result := make([]rune, length)

	for i := 0; i < length; i++ {
		result[i] = strings[rand.Intn(len(strings))]
	}

	return string(result)
}
