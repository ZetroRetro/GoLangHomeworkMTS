package storages

import (
	books "github.com/ZetroRetro/GoLangHomeworkMTS/hw1/book_models"

	"fmt"
)

type MapStorage struct {
	bookByID map[uint64]books.LibraryBook
}

func NewMapStorage() *MapStorage {
	return &MapStorage{bookByID: make(map[uint64]books.LibraryBook, 0)}
}

func (ms *MapStorage) AddBook(book books.LibraryBook) error {
	if valid, err := CheckBook(&book); !valid {
		return fmt.Errorf("invalid book: %w", err)
	}

	ms.bookByID[book.ID] = book
	return nil
}

func (ms *MapStorage) FindBook(id uint64) (*books.LibraryBook, bool) {
	book, ok := ms.bookByID[id]
	return &book, ok
}

func (ms *MapStorage) GiveBook(id uint64, to string) (*books.LibraryBook, error) {
	book, ok := ms.bookByID[id]
	if !ok {
		return nil, fmt.Errorf("storage give error: %w", ErrNonExistentBook)
	}
	if !book.IsOnHand {
		return nil, fmt.Errorf("storage give error: %w", ErrBookUnavailable)
	}

	book.IsOnHand = false
	book.CurrentlyAt = to
	return &book, nil
}

func (ms *MapStorage) ReturnBook(id uint64, from string) error {
	book, ok := ms.bookByID[id]
	if !ok {
		return fmt.Errorf("storage return error: %w", ErrNonExistentBook)
	}

	book.IsOnHand = true
	book.CurrentlyAt = ""
	return nil
}

func (ms *MapStorage) DeleteBook(id uint64) error {
	if _, ok := ms.bookByID[id]; !ok {
		return ErrNonExistentBook
	}
	delete(ms.bookByID, id)
	return nil
}

func (ms *MapStorage) Clear() {
	ms.bookByID = make(map[uint64]books.LibraryBook, 0)
}

func (ms *MapStorage) GetAllBooks() []books.LibraryBook {
	result := make([]books.LibraryBook, 0, len(ms.bookByID))
	for _, value := range ms.bookByID {
		result = append(result, value)
	}
	return result
}
