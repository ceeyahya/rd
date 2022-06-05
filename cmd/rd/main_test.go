package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/ceeyahya/rd"
)

var (
	binName  = "rd"
	fileName = ".books.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building...")

	if runtime.GOOS == "windowns" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "couldn't build %s. stack trace: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running Tests...")
	result := m.Run()

	fmt.Println("Cleaning Up...")

	os.Remove(binName)
	os.Exit(result)
}

func TestCLI(t *testing.T) {
	book := rd.Book{
		Title:  "The Go Programming Language",
		Author: "Brian Kernighan, Alan Donovan",
	}

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(dir, binName)

	t.Run("AddBook", func(t *testing.T) {
		cmd := exec.Command(path, "-title", book.Title, "-author", book.Author)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListBooks", func(t *testing.T) {
		cmd := exec.Command(path, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf(" 1: %s by %s\n", book.Title, book.Author)

		if expected != string(out) {
			t.Errorf("expected %q, instead got %q", expected, string(out))
		}
	})
}
