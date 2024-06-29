package repository

import "github.com/ace-zhaoy/errors"

var (
	ErrNotFound      = errors.NewWithMessage("repository: record not found")
	ErrDuplicatedKey = errors.NewWithMessage("repository: duplicated key not allowed")
)
