package main

import (
	"strings"
	"time"
)

// Transaction -
type Transaction struct {
	date   time.Time
	first  *User
	second *User
	value  string
	tType  string
}

// GetTransactions -
func GetTransactions() (s string) {
	for _, v := range transactions {
		s += strings.Join([]string{v.date.Format(time.RFC3339), v.first.name, v.tType, v.value, "to", v.second.name, "\n"}, " ")
	}
	return s
}

// CreateTransaction -
func CreateTransaction(first, second *User, tType, val string) {
	transactions = append(transactions, &Transaction{date: time.Now(), first: first, second: second, value: val, tType: tType})
}
