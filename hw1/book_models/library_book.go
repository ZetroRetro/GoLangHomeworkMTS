package book

import "fmt"

type LibraryBook struct {
	IDBook
	IsOnHand    bool
	CurrentlyAt string
}

func (book LibraryBook) String() string {
	in_stock := "unavailable"
	if book.IsOnHand {
		in_stock = "available"
	}

	return fmt.Sprintf("#%d (%s): %s", book.IDBook.ID, in_stock, book.BookBase.String())
}
