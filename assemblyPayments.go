package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	money "github.com/Rhymond/go-money"
)

var (
	transactions []*Transaction
	users        []*User
	books        []*Book
)

func init() {
	// Note: The money value is in cents
	// eg. 100 == $1
	// The currency 'type' needs to be consistent if you want to
	// use multiple values together (eg. addition, subtraction... etc)
	users = []*User{
		&User{"Bob", money.New(10000, "AUD")},
		&User{"Kelly", money.New(10000, "AUD")},
	}
	books = []*Book{
		&Book{
			title:   "CrazyTown",
			forSale: true,
			price:   money.New(1000, "AUD"),
			owner:   users[0],
		},
		&Book{
			title:   "The Go Programming Language",
			forSale: false,
			price:   money.New(1200, "AUD"),
			owner:   users[0],
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
			vals, err := GetBankBalance(arg, users)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Print(vals)
			}
		case "books":
			fmt.Print(GetBookList(arg, books, users))
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
   Books - List all the books for sale, or all of the books owned by the given name.
   Exit - Exits this program.`)

}
