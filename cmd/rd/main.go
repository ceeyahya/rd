package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ceeyahya/rd"
)

var booksFileName = ".books.json"

func main() {
	if os.Getenv("BIBDIR") != "" {
		booksFileName = os.Getenv("BIBDIR")
	}

	title := flag.String("title", "", "Book to be added to the bibliography")
	author := flag.String("author", "", "Author of the book to be added to the bibliography")
	list := flag.Bool("list", false, "List the bibliography")
	read := flag.Int("read", 0, "Mark book as read")

	flag.Parse()

	b := &rd.Bibliography{}

	if err := b.GetAllBooks(booksFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(b)

	case *read > 0:
		if err := b.MarkAsFinished(*read); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := b.Save(booksFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *title != "" && *author != "":
		newB := rd.Book{
			Title:  *title,
			Author: *author,
		}
		b.AddBook(newB)
		if err := b.Save(booksFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "The flag(s) provided was not recognized")
		os.Exit(1)
	}
}
