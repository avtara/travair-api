package businesses

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")
	ErrDuplicateData = errors.New("duplicate data")
)