package assembly_test

import (
	"testing"

	money "github.com/Rhymond/go-money"
	assembly "github.com/shanehowearth/assemblyPayments/pkg"
)

func TestMain(m *testing.M) {

	testUsers = []*assembly.User{
		&assembly.User{Name: "TestUserA", BankBal: money.New(10000, "AUD")},
		&assembly.User{Name: "TestUserB", BankBal: money.New(10000, "AUD")},
	}
	testBooks = []*assembly.Book{

		&assembly.Book{Title: "TestBook", ForSale: true, Price: money.New(1000, "AUD"), Owner: testUsers[0]},
		&assembly.Book{Title: "TestBookB", ForSale: true, Price: money.New(1000, "AUD"), Owner: testUsers[1]},
		&assembly.Book{Title: "Not For Sale", ForSale: false, Price: money.New(1000, "AUD"), Owner: testUsers[1]},
	}
	m.Run()
}

var (
	testUsers        []*assembly.User
	testBooks        []*assembly.Book
	testTransactions []*assembly.Transaction
)

func testReset() {
	// Reset user's money
	testUsers[0].BankBal = money.New(10000, "AUD")
	testUsers[1].BankBal = money.New(10000, "AUD")

	// Reset book state
	testBooks[0].Owner = testUsers[0]
	testBooks[0].ForSale = true

	// Clear all transactions generated by tests
	testTransactions = testTransactions[:0]
}