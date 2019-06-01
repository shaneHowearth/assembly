package assembly

import (
	"fmt"
	"strconv"
	"strings"

	money "github.com/Rhymond/go-money"
)

// Book -
type Book struct {
	Title   string
	ForSale bool
	Price   *money.Money
	Owner   *User
}

// GetBook - get a book using a case insensitive search
func GetBook(bookname string, bookList []*Book) (b *Book) {
	for _, bval := range bookList {
		if strings.ToLower(bval.Title) == strings.ToLower(strings.Trim(bookname, " ")) {
			b = bval
			break
		}
	}
	return b
}

// ChangeOwner -
func (b *Book) ChangeOwner(newOwner *User) {
	b.Owner = newOwner
	b.ForSale = false
}

// GetBookList - Get Booklist for the named user, or for everyone
func GetBookList(username string, bookList []*Book, userList []*User) string {
	var bookString string
	if username == "" {
		// All the books that are currently for sale
		for _, book := range bookList {
			if book.ForSale {
				bookString += fmt.Sprintf("%s\n", book.Title)
			}
		}
	} else {
		u, err := FindUser(username, userList)
		if err != nil {
			return "User not found"
		}
		for _, book := range bookList {
			if book.Owner == u {
				bookString += fmt.Sprintf("%s\n", book.Title)
			}
		}
	}
	return bookString
}

// Bid -
func Bid(arg string, bookList []*Book, userList []*User, transactionList []*Transaction) (err error) {
	argList := strings.Split(arg, ",")
	argLength := len(argList)
	if argLength == 3 {
		username := strings.ToLower(argList[0])
		var bidVal float64
		bidVal, err = strconv.ParseFloat(strings.Trim(argList[argLength-1], " "), 64)
		if err != nil {
			return err
		}
		bidVal *= 100 // fix for dollars
		bidAmount := int64(bidVal)
		bookTitle := strings.Join(argList[1:argLength-1], " ")
		wantedBook := GetBook(bookTitle, bookList)
		if wantedBook == nil {
			return fmt.Errorf("No book with that title found")
		}
		u, err := FindUser(username, userList)
		if err != nil {
			return err
		}
		return u.BuyBook(wantedBook, bidAmount, transactionList)
	}

	return fmt.Errorf("Usage: bid(username, bookname, bidAmount)")
}
