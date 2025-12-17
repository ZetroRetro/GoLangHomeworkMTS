package storages

import (
	books "github.com/ZetroRetro/GoLangHomeworkMTS/hw1/book_models"
)

func CheckBook(book *books.LibraryBook) (bool, error) {
	if book.Info.PageCount < 0 {
		return false, ErrPagesUnderflow
	}
	if book.ID == 0 {
		return false, ErrInvalidId
	}
	if book.Info.Title == "" {
		return false, ErrEmptyTitle
	}
	if book.Info.Author == "" {
		return false, ErrNoAuthor
	}

	return true, nil
}
