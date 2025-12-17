package main

import (
	"fmt"
	"log"

	book "github.com/ZetroRetro/GoLangHomeworkMTS/hw1/book_models"
	"github.com/ZetroRetro/GoLangHomeworkMTS/hw1/generator"
	library "github.com/ZetroRetro/GoLangHomeworkMTS/hw1/libriary"
	"github.com/ZetroRetro/GoLangHomeworkMTS/hw1/storages"
)

func main() {
	books := []book.BookBase{
		{
			Info: book.BookInfoObj{
				Title:       "Война и мир",
				Author:      "Лев Толстой",
				PageCount:   1225,
				PublishYear: 1869,
			},
			Text: "Полный текст книги...",
		},
		{
			Info: book.BookInfoObj{
				Title:       "Преступление и наказание",
				Author:      "Фёдор Достоевский",
				PageCount:   672,
				PublishYear: 1866,
			},
			Text: "Полный текст книги...",
		},
		{
			Info: book.BookInfoObj{
				Title:       "Мастер и Маргарита",
				Author:      "Михаил Булгаков",
				PageCount:   480,
				PublishYear: 1967,
			},
			Text: "Полный текст книги...",
		},
		{
			Info: book.BookInfoObj{
				Title:       "1984",
				Author:      "Джордж Оруэлл",
				PageCount:   328,
				PublishYear: 1949,
			},
			Text: "Полный текст книги...",
		},
	}

	storage := storages.NewSliceStorage()
	idGen := generator.SequenceGenerator(1)
	lib := library.NewLibrary(storage, idGen)

	fmt.Println("=== Загрузка книг в библиотеку ===")
	for _, b := range books {
		addedBook, err := lib.AddBook(b)
		if err != nil {
			log.Printf("Ошибка добавления книги '%s': %v", b.Info.Title, err)
		} else {
			fmt.Printf("Добавлена: %s\n", addedBook.String())
		}
	}

	fmt.Println("\n=== Поиск книг в библиотеке ===")

	foundBook1, ok1 := lib.FindBook("Война и мир — Лев Толстой — 1869")
	if ok1 {
		fmt.Printf("Найдена книга 1: %s\n", foundBook1.String())
	} else {
		fmt.Println("Книга 1 не найдена")
	}

	foundBook2, ok2 := lib.FindBook("1984 — Джордж Оруэлл — 1949")
	if ok2 {
		fmt.Printf("Найдена книга 2: %s\n", foundBook2.String())
	} else {
		fmt.Println("Книга 2 не найдена")
	}

	fmt.Println("\n=== Замена генератора ID ===")
	newIdGen := generator.FromAddressGenerator()
	if err := lib.ReplaceIDGenerator(newIdGen); err != nil {
		log.Printf("Ошибка замены генератора: %v", err)
	} else {
		fmt.Println("Генератор ID успешно заменен")
	}

	fmt.Println("\n=== Добавление книги с новым генератором ID ===")
	newBook := book.BookBase{
		Info: book.BookInfoObj{
			Title:       "Гарри Поттер и философский камень",
			Author:      "Джоан Роулинг",
			PageCount:   432,
			PublishYear: 1997,
		},
		Text: "Полный текст книги...",
	}

	addedNewBook, err := lib.AddBook(newBook)
	if err != nil {
		log.Printf("Ошибка добавления книги: %v", err)
	} else {
		fmt.Printf("Добавлена новая книга: %s\n", addedNewBook.String())
	}

	fmt.Println("\n=== Поиск новой книги ===")
	foundNewBook, ok3 := lib.FindBook("Гарри Поттер и философский камень — Джоан Роулинг — 1997")
	if ok3 {
		fmt.Printf("Найдена новая книга: %s\n", foundNewBook.String())
	} else {
		fmt.Println("Новая книга не найдена")
	}

	fmt.Println("\n=== Замена хранилища ===")
	newStorage := storages.NewMapStorage()
	if err := lib.RebuildStorage(newStorage); err != nil {
		log.Printf("Ошибка замены хранилища: %v", err)
	} else {
		fmt.Println("Хранилище успешно заменено")
	}

	fmt.Println("\n=== Добавление книг в новое хранилище ===")
	additionalBooks := []book.BookBase{
		{
			Info: book.BookInfoObj{
				Title:       "Убить пересмешника",
				Author:      "Харпер Ли",
				PageCount:   384,
				PublishYear: 1960,
			},
			Text: "Полный текст книги...",
		},
		{
			Info: book.BookInfoObj{
				Title:       "Великий Гэтсби",
				Author:      "Фрэнсис Скотт Фицджеральд",
				PageCount:   218,
				PublishYear: 1925,
			},
			Text: "Полный текст книги...",
		},
	}

	if err := lib.AddBooks(additionalBooks); err != nil {
		log.Printf("Ошибка добавления дополнительных книг: %v", err)
	} else {
		fmt.Println("Дополнительные книги добавлены")
	}

	fmt.Println("\n=== Поиск книг в новом хранилище ===")

	foundExisting, ok4 := lib.FindBook("Мастер и Маргарита — Михаил Булгаков — 1967")
	if ok4 {
		fmt.Printf("Найдена существующая книга: %s\n", foundExisting.String())
	} else {
		fmt.Println("Существующая книга не найдена")
	}

	foundAdditional, ok5 := lib.FindBook("Убить пересмешника — Харпер Ли — 1960")
	if ok5 {
		fmt.Printf("Найдена дополнительная книга: %s\n", foundAdditional.String())
	} else {
		fmt.Println("Дополнительная книга не найдена")
	}

	fmt.Println("\n=== Работа библиотеки завершена ===")
}
