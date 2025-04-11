package middleware

import (
	"encoding/base64"
	"strings"
)

func isValidBasicAuth(authToken string) bool {

	byteToken, _ := base64.StdEncoding.DecodeString(strings.TrimPrefix(authToken, "Basic "))
	key := strings.Split(string(byteToken), ":")
	if len(key) < 2 {
		return false
	}
	username, password := key[0], key[1]

	if username == "nur" && password == "r1zk1" {
		return true
	}
	return false
}
