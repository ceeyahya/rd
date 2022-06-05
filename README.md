# rd

rd is utility CLI application that allow me to easily manage my technical reading bibliography. This started as a Google Sheet and ended being hard to manage so I decided to turn it
into a CLI to make it more convinient.

## Usage

The bibiliography is stored in a json file that by default is created in the same folder as the binary and is called `.books.json` however you can change that behavior through an environment
variable `BIBDIR` using the following command:

- Unix

```console
user@user:~$ export BIBDIR=/path/to/bibliography
```

- Windows

```console
user@user:~$ set BIBDIR=/path/to/bibliography
```

### Adding a new book to the bibliography:

Adding a book to the bibliography simply appends a new entry to the json file using:

```console
user@user:~$ rd -title "The Go Programming Language" -author "Brian Kernighan, Alan Donovan"
```

### List all the books

Listing all the books displays them nicely with their ids and shows whether they've been read or not using:

```console
user@user:~$ rd -list
  1: The Go Programming Language by Brian Kernighan, Alan Donovan
  2: Advanced Algorithms and Data Structures by Marcello La Rocca
  3: The Ultimate Go Notebook by Ardan Labs
  4: Rust in Action by Tim McNamara
  5: Zero to Production by Luca Palmieri
  6: Rust for Rustaceans by Jon Gjengset
  7: Programming Rust by Jim Blandy, Jason Orendorff, Leonora F. S. Tindall
```

### Mark a book as completed

It is possible to mark a certain book as read or completed via its id by using:

```console
user@user:~$ rd -complete <id>
```

the output of the `-list` command is then updated to look something like this:

```console
user@user:~$ rd -list
✓ 1: The Go Programming Language by Brian Kernighan, Alan Donovan
  2: Advanced Algorithms and Data Structures by Marcello La Rocca
✓ 3: The Ultimate Go Notebook by Ardan Labs
  4: Rust in Action by Tim McNamara
  5: Zero to Production by Luca Palmieri
  6: Rust for Rustaceans by Jon Gjengset
  7: Programming Rust by Jim Blandy, Jason Orendorff, Leonora F. S. Tindall
```
