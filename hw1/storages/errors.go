package storages

import "errors"

var (
	// Book errors
	ErrPagesUnderflow = errors.New("pages must be non-negative")
	ErrInvalidId      = errors.New("id must be a positive number")
	ErrEmptyTitle     = errors.New("non empty title is required")
	ErrNoAuthor       = errors.New("author is required")

	// Storage errors
	ErrNonExistentBook = errors.New("there no such book")
	ErrBookUnavailable = errors.New("book is temporary unavailable")
)
