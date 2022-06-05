package rd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Book struct {
	Title      string
	Author     string
	Read       bool
	CreatedAt  time.Time
	FinishedAt time.Time
}

type Bibliography []Book

func (b *Bibliography) String() string {
	formatted := ""
	for i, b := range *b {
		prefix := "  "
		if b.Read {
			prefix = "✓ "
		}
		formatted += fmt.Sprintf("%s%d: %s by %s\n", prefix, i+1, b.Title, b.Author)
	}
	return formatted
}

func (b *Bibliography) AddBook(book Book) {
	newBook := Book{
		Title:      book.Title,
		Author:     book.Author,
		Read:       book.Read,
		CreatedAt:  time.Now(),
		FinishedAt: time.Time{},
	}

	*b = append(*b, newBook)
}

func (b *Bibliography) MarkAsFinished(i int) error {
	biblio := *b
	if i <= 0 || i > len(biblio) {
		return fmt.Errorf("Book number %d does not exist", i)
	}

	biblio[i-1].Read = true
	biblio[i-1].FinishedAt = time.Now()

	return nil
}

func (b *Bibliography) DeleteBook(i int) error {
	biblio := *b
	if i <= 0 || i > len(biblio) {
		return fmt.Errorf("Book number %d does not exist", i)
	}

	*b = append(biblio[:i-1], biblio[i:]...)
	return nil
}

func (b *Bibliography) Save(filename string) error {
	json, err := json.Marshal(b)
	if err != nil {
		return fmt.Errorf("error encoutered while saving to %s. Stack Trace: %s", filename, err)
	}
	return os.WriteFile(filename, json, 0644)
}

func (b *Bibliography) GetAllBooks(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("error encoutered while saving to %s. Stack Trace: %s", filename, err)
	}

	if len(file) == 0 {
		return fmt.Errorf("the requested file is empty")
	}

	return json.Unmarshal(file, b)
}
