package errno

import "net/http"

var (
	// ErrBind gin bind error
	ErrBind = &Err{
		ResultCode: CodeValidError,
		Code:       "BIND_ERR",
		Message:    "Error occurred while binding the request body to the struct.",
		StatusCode: http.StatusBadRequest}
	// ErrValidation gin validation error
	ErrValidation = &Err{
		ResultCode: CodeValidError,
		Code:       "VALIDATION_ERR",
		Message:    "Validation failed.",
		StatusCode: http.StatusBadRequest}
	// ErrEncrypt error
	ErrEncrypt = &Err{
		ResultCode: CodeValidError,
		Code:       "ENCRYPT_ERR",
		Message:    "Error occurred while encrypting the user password.",
		StatusCode: http.StatusInternalServerError}
	// ErrDatabase error
	ErrDatabase = &Err{
		ResultCode: CodeServerError,
		Code:       "DB_ERR",
		Message:    "Database error.",
		StatusCode: http.StatusInternalServerError}
	// ErrNotFound error
	ErrNotFound = &Err{
		ResultCode: CodeServerError,
		Code:       "NOT_FOUND",
		Message:    "404 not found.",
		StatusCode: http.StatusNotFound}
	// ErrUserNotFound error
	ErrUserNotFound = &Err{
		ResultCode: CodeServerError,
		Code:       "USER_NOT_FOUND",
		Message:    "The user was not found.",
		StatusCode: http.StatusNotFound}
	// ErrTimeout error
	ErrTimeout = &Err{
		ResultCode: CodeServerError,
		Code:       "TIMEOUT",
		Message:    "Server timed out waiting for the request",
		StatusCode: http.StatusRequestTimeout}
	// ErrTokenInvalid error
	ErrTokenInvalid = &Err{
		ResultCode: CodeAuthError,
		Code:       "TOKEN_INVALID",
		Message:    "The token was invalid.",
		StatusCode: http.StatusForbidden}
	// ErrPasswordIncorrect error
	ErrPasswordIncorrect = &Err{
		ResultCode: CodeAuthError,
		Code:       "PASSWORD_INCORRECT",
		Message:    "The password was incorrect.",
		StatusCode: http.StatusForbidden}
	// ErrToken error
	ErrToken = &Err{
		ResultCode: CodeAuthError,
		Code:       "TOKEN_ERR",
		Message:    "Error occurred while signing the JSON web token.",
		StatusCode: http.StatusUnauthorized}
)
