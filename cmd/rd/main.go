package main

import (
	"fmt"
	"os"

	"github.com/ceeyahya/rd"
)

const booksFileName = ".books.json"

func main() {
	b := &rd.Bibliography{}

	switch {
	case len(os.Args) == 1:
		if err := b.GetAllBooks(booksFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, item := range *b {
			fmt.Println(item.Title)
		}

	default:
		title := os.Args[1]
		author := os.Args[2]
		book := rd.Book{
			Title:  title,
			Author: author,
		}

		b.AddBook(book)
		if err := b.Save(booksFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
