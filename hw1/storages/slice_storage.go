package storages

import (
	"fmt"

	books "github.com/ZetroRetro/GoLangHomeworkMTS/hw1/book_models"
)

var emptyBookSpace = books.LibraryBook{}

type SliceStorage struct {
	items []books.LibraryBook
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{items: make([]books.LibraryBook, 0)}
}

func (ss *SliceStorage) AddBook(book books.LibraryBook) error {
	if valid, err := CheckBook(&book); valid {
		return fmt.Errorf("invalid book: %w", err)
	}

	for i, item := range ss.items {
		if item.ID == book.ID || item.ID == 0 {
			ss.items[i] = book
			return nil
		}
	}

	ss.items = append(ss.items, book)
	return nil
}

func (ss *SliceStorage) FindBook(id uint64) (*books.LibraryBook, bool) {
	for _, item := range ss.items {
		if item.ID == id {
			return &item, true
		}
	}
	return nil, false
}

func (ss *SliceStorage) GiveBook(id uint64, to string) (*books.LibraryBook, error) {
	book, found := ss.FindBook(id)
	if !found {
		return nil, fmt.Errorf("storage give error: %w", ErrNonExistentBook)
	}

	if !book.IsOnHand {
		return nil, fmt.Errorf("storage give error: %w", ErrBookUnavailable)
	}

	book.IsOnHand = true
	book.CurrentlyAt = to

	return book, nil
}

func (ss *SliceStorage) ReturnBook(id uint64, from string) error {
	for i, item := range ss.items {
		if item.ID == id && item.CurrentlyAt == from {
			ss.items[i].CurrentlyAt = ""
			ss.items[i].IsOnHand = true
			return nil
		}
	}

	return fmt.Errorf("storage return error: %w", ErrNonExistentBook)
}

func (ss *SliceStorage) DeleteBook(id uint64) error {
	for index := range ss.items {
		if ss.items[index].ID == id {
			ss.items = append(ss.items[:index], ss.items[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("storage delete error: %w", ErrNonExistentBook)
}

func (ss *SliceStorage) Clear() {
	ss.items = make([]books.LibraryBook, 0)
}

func (ss *SliceStorage) GetAllBooks() []books.LibraryBook {
	return append([]books.LibraryBook(nil), ss.items...)
}
