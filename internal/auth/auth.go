package auth

import (
	"errors"
	"net/http"
	"strings"
)

const Err_No_Auth_Header_Included string = "no authorization header included"
const Err_Malformed_Auth_Header string = "malformed authorization header"

const Get_Authorization string = "Authorization"
const Get_Api_Key string = "ApiKey"

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get(Get_Authorization)
	if authHeader == "" {
		return "", errors.New(Err_No_Auth_Header_Included)
	}

	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != Get_Api_Key {
		return "", errors.New(Err_Malformed_Auth_Header)
	}

	return splitAuth[1], nil
}
