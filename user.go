package main

import (
	"fmt"
	"strconv"
	"strings"

	money "github.com/Rhymond/go-money"
)

// User -
type User struct {
	name    string
	bankBal *money.Money
}

// FindUser - Find user by name (case insensitive)
func FindUser(username string, userList []*User) (*User, error) {
	for _, v := range userList {
		if strings.ToLower(v.name) == strings.Trim(strings.ToLower(username), " ") {
			return v, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}

// GetBankBalance - Get bank balance of named user, or for everyone
func GetBankBalance(username string, userList []*User) (string, error) {
	if username == "" {
		var everyone string
		for _, v := range users {
			everyone += fmt.Sprintf("%s - %d\n", v.name, v.bankBal.Amount()/100)
		}
		return everyone, nil
	}
	foundUser, err := FindUser(username, userList)
	if err != nil {
		return "", fmt.Errorf("Not found")
	}
	return strconv.FormatInt(foundUser.bankBal.Amount()/100, 10) + "\n", nil
}
