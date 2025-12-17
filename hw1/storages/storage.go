package storages

import books "github.com/ZetroRetro/GoLangHomeworkMTS/hw1/book_models"

type Storage interface {
	AddBook(book books.LibraryBook) error
	FindBook(id uint64) (*books.LibraryBook, bool)
	GiveBook(id uint64, to string) (*books.LibraryBook, error)
	ReturnBook(id uint64, from string) error
	DeleteBook(id uint64) error
	Clear()
	GetAllBooks() []books.LibraryBook
}
