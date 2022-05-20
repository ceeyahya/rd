package rd_test

import (
	"testing"

	"github.com/ceeyahya/rd"
)

func TestAddBook(t *testing.T) {
	b := rd.Bibliography{}
	item := rd.Book{
		Title:  "The Go Programming Language",
		Author: "Brian Kernighan, Alan Donovan",
	}

	b.AddBook(item)

	if b[0].Title != item.Title {
		t.Errorf("expected %q but got %q", b[0].Title, item.Title)
	}
}

func TestMarkAsFinished(t *testing.T) {
	b := rd.Bibliography{}
	item := rd.Book{
		Title:  "The Go Programming Language",
		Author: "Brian Kernighan, Alan Donovan",
	}

	b.AddBook(item)

	if b[0].Title != item.Title {
		t.Errorf("expected %q but got %q", b[0].Title, item.Title)
	}

	if b[0].Read {
		t.Errorf("newly added books should be marked as not finished")
	}

	b[0].Read = true

	if !b[0].Read {
		t.Errorf("%q should be marked as finished", b[0].Title)
	}
}

func TestDeleteBook(t *testing.T) {
	b := rd.Bibliography{}
	items := []rd.Book{
		{
			Title:  "The Go Programming Language",
			Author: "Brian Kernighan, Alan Donovan",
		},
		{
			Title:  "Advanced Algorithms and Data Structures",
			Author: "Marcello La Rocca",
		},
		{
			Title:  "Ultimate Go Notebook",
			Author: "William Kennedy",
		},
	}

	for _, v := range items {
		b.AddBook(v)
	}

	if b[0].Title != items[0].Title {
		t.Errorf("expected %q, but got %q", items[0].Title, b[0].Title)
	}

	b.DeleteBook(1)

	if len(b) != 2 {
		t.Errorf("expected the length of the bibliography to be %d, instead got %d", 2, len(b))
	}

	if b[0].Title != items[1].Title {
		t.Errorf("expected %q, but got %q", items[0].Title, b[0].Title)
	}
}
