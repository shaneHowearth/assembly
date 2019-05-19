package main

import (
	"fmt"

	money "github.com/Rhymond/go-money"
)

// Book -
type Book struct {
	title   string
	forSale bool
	price   *money.Money
	owner   *User
}

// GetBookList - Get Booklist for the named user, or for everyone
func GetBookList(username string, bookList []*Book, userList []*User) string {
	var bookString string
	if username == "" {
		// All the books that are currently for sale
		for _, book := range bookList {
			if book.forSale {
				bookString += fmt.Sprintf("%s\n", book.title)
			}
		}
	} else {
		u, err := FindUser(username, userList)
		if err != nil {
			return "User not found"
		}
		for _, book := range bookList {
			if book.owner == u {
				bookString += fmt.Sprintf("%s\n", book.title)
			}
		}
	}
	return bookString
}
