package assembly

import (
	"fmt"
	"strconv"
	"strings"

	money "github.com/Rhymond/go-money"
)

// User -
type User struct {
	Name    string
	BankBal *money.Money
}

// PayMoney - pay money from this user instance' account to the nnamed user
func (u *User) PayMoney(payee *User, amount *money.Money, transactionList *[]*Transaction) (err error) {
	u.BankBal, err = u.BankBal.Subtract(amount)
	if err != nil {
		return err
	}
	payee.BankBal, err = payee.BankBal.Add(amount)
	if err != nil {
		// Roll back the subtraction
		u.BankBal, err = u.BankBal.Add(amount)
		return err
	}
	CreateTransaction(u, payee, "paid", strconv.FormatFloat(float64(amount.Amount()/100), 'f', 2, 64), transactionList)
	return nil
}

// BuyBook -
func (u *User) BuyBook(wantedBook *Book, bidValue int64, transactionList *[]*Transaction) (err error) {
	if wantedBook == nil {
		fmt.Println("No book found by that name, are you sure you typed it in correctly?")
		return
	}

	if wantedBook.Owner != u && wantedBook.ForSale {
		// check bidder has the money to spend
		if u.BankBal.Amount() < bidValue {
			return fmt.Errorf("Sorry, you do not have enough available funds to make that bid")
		}
		// The bid must be equal to or higher than the asking price
		if bidValue >= (wantedBook.Price.Amount()) {

			// pay monies
			err = u.PayMoney(wantedBook.Owner, money.New(bidValue, "AUD"), transactionList)
			if err != nil {
				return fmt.Errorf("Unable to make payment with error: %v", err)
			}
			// transfer ownership
			CreateTransaction(wantedBook.Owner, u, "sold", wantedBook.Title, transactionList)
			wantedBook.ChangeOwner(u)
			fmt.Printf("Congratulations, your bid was successful. You are now the new owner of %s\n", wantedBook.Title)
		} else {
			fmt.Println("Sorry, your bid was unsuccessful")
		}
	} else {
		return fmt.Errorf("%s cannot be sold to %s", wantedBook.Title, u.Name)
	}

	return
}

// FindUser - Find user by name (case insensitive)
func FindUser(username string, userList []*User) (*User, error) {
	for _, v := range userList {
		if strings.ToLower(v.Name) == strings.Trim(strings.ToLower(username), " ") {
			return v, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}

// GetBankBalance - Get bank balance of named user, or for everyone
func GetBankBalance(username string, userList []*User) (string, error) {
	if username == "" {
		var everyone string
		for _, v := range userList {
			everyone += fmt.Sprintf("%s - %d\n", v.Name, v.BankBal.Amount()/100)
		}
		return everyone, nil
	}
	foundUser, err := FindUser(username, userList)
	if err != nil {
		return "", fmt.Errorf("Not found")
	}
	return strconv.FormatInt(foundUser.BankBal.Amount()/100, 10) + "\n", nil
}
