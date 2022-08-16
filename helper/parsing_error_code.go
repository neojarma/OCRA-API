package helper

import "strings"

func IsInvalidForeignKey(err error) bool {
	errStr := strings.Split(err.Error(), " ")
	return errStr[1] == "1452:"
}
