package book

import (
	"fmt"
	"strings"
)

type BookBase struct {
	Info BookInfoObj
	Text string
}

func (book BookBase) String() string {
	parts := []string{
		strings.TrimSpace(book.Info.Title),
		strings.TrimSpace(book.Info.Author),
	}

	if book.Info.PublishYear != 0 {
		parts = append(parts, fmt.Sprintf("%d", book.Info.PublishYear))
	}

	return strings.Join(parts, " — ")
}
