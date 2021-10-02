package businesses

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")
	ErrEmailDuplicate = errors.New("email is already taken")
)