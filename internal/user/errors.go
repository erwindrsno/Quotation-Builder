package user

import "errors"

var (
	errInvalidCredentials = errors.New("invalid username or password.")
	errInternalError      = errors.New("internal error occured.")
)
