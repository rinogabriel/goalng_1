package main

import (
	"errors"
	"fmt"
	"os"
)

type Book struct {
	Title  string
	Author string
	ISBN   string
}

func displayBook(b Book) {
	fmt.Printf("Title: %s\nAuthor: %s\nISBN: %s\n", b.Title, b.Author, b.ISBN)
}

func validateISBN(isbn string) error {
	if len(isbn) != 13 {
		return errors.New("ISBN must be 13 characters long")
	}
	return nil
}

func createBook() (Book, error) {
	var title, author, isbn string

	fmt.Print("Enter Book Title: ")
	fmt.Scanln(&title)

	fmt.Print("Enter Author: ")
	fmt.Scanln(&author)

	fmt.Print("Enter ISBN: ")
	fmt.Scanln(&isbn)

	err := validateISBN(isbn)
	if err != nil {
		return Book{}, err
	}

	return Book{Title: title, Author: author, ISBN: isbn}, nil
}

func main() {
	var library []Book

	for {
		fmt.Println("\n1. Add a Book")
		fmt.Println("2. Display Book Details")
		fmt.Println("3. Exit")
		fmt.Print("Select an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			book, err := createBook()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			library = append(library, book)
			fmt.Println("Book added to the library!")

		case 2:
			fmt.Println("\n--- Library ---")
			for _, book := range library {
				displayBook(book)
				fmt.Println("---")
			}

		case 3:
			fmt.Println("Exiting the program.")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
