package auth

import (
	"errors"
	"net/http"
	"strings"
)

const ErrNoAuthHeaderIncluded = "no authorization header included"
const ErrMalformedAuthHeader = "malformed authorization header"

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New(ErrNoAuthHeaderIncluded)
	}

	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New(ErrMalformedAuthHeader)
	}

	return splitAuth[1], nil
}
