package book

import "fmt"

type IDBook struct {
	ID uint64
	BookBase
}

func (book IDBook) String() string {
	return fmt.Sprintf("#%d: %s", book.ID, book.BookBase.String())
}
