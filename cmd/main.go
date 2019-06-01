package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	money "github.com/Rhymond/go-money"
	assembly "github.com/shanehowearth/assemblyPayments/pkg"
)

var (
	transactions []*assembly.Transaction
	users        []*assembly.User
	books        []*assembly.Book
)

func init() {
	// Note: The money value is in cents
	// eg. 100 == $1
	// The currency 'type' needs to be consistent if you want to
	// use multiple values together (eg. addition, subtraction... etc)
	users = []*assembly.User{
		&assembly.User{Name: "Bob", BankBal: money.New(10000, "AUD")},
		&assembly.User{Name: "Kelly", BankBal: money.New(10000, "AUD")},
	}
	books = []*assembly.Book{
		&assembly.Book{
			Title:   "CrazyTown",
			ForSale: true,
			Price:   money.New(1000, "AUD"),
			Owner:   users[0],
		},
		&assembly.Book{
			Title:   "The Go Programming Language",
			ForSale: false,
			Price:   money.New(1200, "AUD"),
			Owner:   users[0],
		},
	}
}

func main() {
	ui()
}

func ui() {
	scanner := bufio.NewScanner(os.Stdin)
	// arguments can be inside brackets, or just spaces
	r1, _ := regexp.Compile("\\(| ")
	r2, _ := regexp.Compile(`[)].*`)
	// Initial Prompt
	fmt.Println("Simple book buying app, Type 'help' for a list of commands.")
	fmt.Print("> ")
	// Infinite loop waiting for input
	for scanner.Scan() {
		c := r1.Split(scanner.Text(), -1)
		// Anything after the command is deemed an argument
		var arg string
		if len(c) > 1 {
			arg = string(r2.ReplaceAll([]byte(strings.Join(c[1:], " ")), []byte("")))
		}
		switch com := strings.ToLower(c[0]); com {
		case "help":
			Help()
		case "bal":
			vals, err := assembly.GetBankBalance(arg, users)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Print(vals)
			}
		case "bid":
			err := assembly.Bid(arg, books, users, transactions)
			if err != nil {
				fmt.Println(err)
			}
		case "books":
			fmt.Print(assembly.GetBookList(arg, books, users))
		case "exit":
			os.Exit(0)
		}
		fmt.Print("> ")
	}

	if scanner.Err() != nil {
		os.Exit(1)
	}
}

// Help -
func Help() {
	fmt.Println(`
Available commands:
   Help - Prints this message.
   Bal - Gets the current bank balances.
       If a name is given then only the bank balance for that person.
   Bid - Bid for a book.
       Give the name of the person bidding, the name of the book they are bidding for, and the amount that they are bidding. If successful the book will change ownership, and the money paid from the bidder's wallet to the seller's.
   Books - List all the books for sale, or all of the books owned by the given name.
   Exit - Exits this program.`)

}
