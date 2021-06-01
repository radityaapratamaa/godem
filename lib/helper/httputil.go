package helper

import (
	"net/http"
	"strings"
)

func GetBearerToken(r *http.Request) string {
	var token string
	if authHeader := r.Header.Get("Authorization"); authHeader != "" {
		token = strings.Replace(authHeader, "Bearer ", "", 1)
	}
	return token
}
