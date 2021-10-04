package businesses

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")
	ErrEmailDuplicate = errors.New("email is already taken")
	ErrAccountActivated = errors.New("the account has been activated")
	ErrAccountNotFound = errors.New("account not found")
	ErrAccountUnactivated = errors.New("account has not been activated")
	ErrInvalidCredential = errors.New("email or password does not match")
)