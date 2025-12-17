package library

import (
	"fmt"
	"strings"

	"github.com/ZetroRetro/GoLangHomeworkMTS/hw1/storages"

	books "github.com/ZetroRetro/GoLangHomeworkMTS/hw1/book_models"
)

var kMAX_ATTEMPTS = 10000

type idGenerator func(...any) uint64

type Library struct {
	storage storages.Storage
	idGen   idGenerator
	nameIDs map[string]uint64
}

func NewLibrary(storage storages.Storage, idGenerator idGenerator) *Library {
	return &Library{
		storage: storage,
		idGen:   idGenerator,
		nameIDs: make(map[string]uint64),
	}
}

func (l *Library) ReplaceIDGenerator(idGenerator idGenerator) error {
	if idGenerator == nil {
		return fmt.Errorf("library generator error: %w", ErrMissingIDGenerator)
	}
	l.idGen = idGenerator
	return nil
}

func (l *Library) RebuildStorage(newStorage storages.Storage) error {
	if newStorage == nil {
		return fmt.Errorf("library storage error: %w", ErrMissingStorage)
	}

	for _, identifiedBook := range l.storage.GetAllBooks() {
		if err := newStorage.AddBook(identifiedBook); err != nil {
			return fmt.Errorf("migrate book id=%d error: %w", identifiedBook.ID, err)
		}
	}

	l.storage.Clear()
	l.storage = newStorage
	l.reIndex()
	return nil
}

func (l *Library) AddBook(new_book books.BookBase) (books.LibraryBook, error) {
	name := new_book.String()
	id, err := l.generateUniqueID(name)
	if err != nil {
		return books.LibraryBook{}, err
	}

	lib_book := books.LibraryBook{
		IDBook: books.IDBook{
			BookBase: new_book,
			ID:       id},
		IsOnHand:    true,
		CurrentlyAt: "",
	}
	if err := l.storage.AddBook(lib_book); err != nil {
		return books.LibraryBook{}, err
	}
	l.nameIDs[name] = id
	return lib_book, nil
}

func (l *Library) AddBooks(books []books.BookBase) error {
	for _, bookToAdd := range books {
		if _, err := l.AddBook(bookToAdd); err != nil {
			return err
		}
	}
	return nil
}

func (l *Library) FindBook(title string) (books.BookBase, bool) {
	id, hasTitle := l.nameIDs[strings.TrimSpace(title)]
	if !hasTitle {
		return books.BookBase{}, false
	}

	book, found := l.storage.FindBook(id)
	return book.BookBase, found
}

func (l *Library) DeleteBook(name string) error {
	id, ok := l.nameIDs[name]
	if !ok {
		return fmt.Errorf("library deletion: %w", ErrNoBook)
	}

	if err := l.storage.DeleteBook(id); err != nil {
		return fmt.Errorf("library deletion: %w", err)
	}

	delete(l.nameIDs, name)
	return nil
}

func (l *Library) reIndex() {
	l.nameIDs = make(map[string]uint64)

	for _, book := range l.storage.GetAllBooks() {
		name := strings.TrimSpace(book.BookBase.String())
		l.nameIDs[name] = book.ID
	}
}

func (l *Library) generateUniqueID(opts ...any) (uint64, error) {
	for range kMAX_ATTEMPTS {
		id := l.idGen(opts...)
		if id == 0 {
			continue
		}
		if _, exists := l.storage.FindBook(id); exists {
			continue
		}
		return id, nil
	}
	return 0, ErrIdGenerate
}
