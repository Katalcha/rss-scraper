package auth

import (
	"errors"
	"net/http"
	"strings"
)

const Err_No_Auth_Header_Included = "no authorization header included"
const Err_Malformed_Auth_Header = "malformed authorization header"

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New(Err_No_Auth_Header_Included)
	}

	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New(Err_Malformed_Auth_Header)
	}

	return splitAuth[1], nil
}
