package library

import (
	"errors"
)

var (
	ErrMissingTitle       = errors.New("title is required")
	ErrTitleAlreadyExists = errors.New("book with this title already exists")
	ErrNoBook             = errors.New("such book does not exist")
	ErrMissingIDGenerator = errors.New("id generator is nil")
	ErrMissingStorage     = errors.New("storage is nil")
	ErrIdGenerate         = errors.New("failed to generate unique id")
)
