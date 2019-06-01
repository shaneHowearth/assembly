package assembly

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
func GetTransactions(transactionList []*Transaction) (s string) {
	for _, v := range transactionList {
		s += strings.Join([]string{v.date.Format(time.RFC3339), v.first.Name, v.tType, v.value, "to", v.second.Name, "\n"}, " ")
	}
	return s
}

// CreateTransaction -
func CreateTransaction(first, second *User, tType, val string, transactionList []*Transaction) {
	transactionList = append(transactionList, &Transaction{date: time.Now(), first: first, second: second, value: val, tType: tType})
}
