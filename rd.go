package rd

import (
	"fmt"
	"time"
)

type Book struct {
	title      string
	author     string
	read       bool
	createdAt  time.Time
	finishedAt time.Time
}

type Bibliography []Book

func (b *Bibliography) AddBook(book Book) {
	newBook := Book{
		title:      book.title,
		author:     book.author,
		read:       book.read,
		createdAt:  time.Now(),
		finishedAt: time.Time{},
	}

	*b = append(*b, newBook)
}

func (b *Bibliography) MarkAsFinished(i int) error {
	biblio := *b
	if i <= 0 || i > len(biblio) {
		return fmt.Errorf("Book number %d does not exist", i)
	}

	biblio[i-1].read = true
	biblio[i-1].finishedAt = time.Now()

	return nil
}
