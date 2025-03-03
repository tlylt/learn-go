package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

type Library struct {
	books map[string]*Book
}

func (l *Library) AddBook(title, author string) {
	l.books[title] = &Book{Title: title, Author: author}
}

func (l *Library) FindBook(title string) *Book {
	if book, ok := l.books[title]; !ok {
		return nil
	} else {
		return book
	}
}

func (l *Library) GetBook(title string) (*Book, error) {
	if book, ok := l.books[title]; !ok {
		return nil, fmt.Errorf("Book not found")
	} else {
		return book, nil
	}
}

func main() {
	library := Library{books: make(map[string]*Book)}
	library.AddBook("Your Code as a Crime Scene", "Adam Tornhill")

	// exist and well
	book := library.FindBook("Your Code as a Crime Scene")

	fmt.Printf("Book: %s, Author: %s\n", book.Title, book.Author)

	// not exist and panic
	book = library.FindBook("The Phoenix Project")

	fmt.Printf("Book: %s, Author: %s\n", book.Title, book.Author)
}
