package assembly

import (
	"strings"
	"time"
)

// Transaction -
type Transaction struct {
	Date   time.Time
	First  *User
	Second *User
	Value  string
	TType  string
}

// GetTransactions -
func GetTransactions(transactionList *[]*Transaction) (s string) {
	for _, v := range *transactionList {
		s += strings.Join([]string{v.Date.Format(time.RFC3339), v.First.Name, v.TType, v.Value, "to", v.Second.Name, "\n"}, " ")
	}
	return s
}

// CreateTransaction -
func CreateTransaction(first, second *User, tType, val string, transactionList *[]*Transaction) {
	*transactionList = append(*transactionList, &Transaction{Date: time.Now(), First: first, Second: second, Value: val, TType: tType})
}
