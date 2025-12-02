package helper

import (
	"net/http"

	"github.com/salmantaghooni/golang-car-web-api/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{
	// OTP
	service_errors.OTPExists:   http.StatusConflict,
	service_errors.OTPUsed:     http.StatusConflict,
	service_errors.OTPNotValid: http.StatusBadRequest,

	// User
	service_errors.EmailExists:      http.StatusConflict,
	service_errors.UsernameExists:   http.StatusConflict,
	service_errors.RecordNotFound:   http.StatusBadRequest,
	service_errors.PermissionDenied: 403,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
