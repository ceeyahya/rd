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
		t.Errorf("Expected %q but got %q", b[0].Title, item.Title)
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
		t.Errorf("Expected %q but got %q", b[0].Title, item.Title)
	}

	if b[0].Read {
		t.Errorf("Newly added books should be marked as not finished")
	}

	b[0].Read = true

	if !b[0].Read {
		t.Errorf("%q should be marked as finished", b[0].Title)
	}
}
