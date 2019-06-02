package assembly_test

import (
	"errors"
	"fmt"
	"testing"

	money "github.com/Rhymond/go-money"
	assembly "assemblyPayments/pkg"
	"github.com/stretchr/testify/assert"
)

// TestGetBankBalance -
func TestGetBankBalance(t *testing.T) {
	testcases := map[string]struct {
		username string
		response string
		err      error
	}{
		"TestUserA": {
			username: "TestUserA",
			response: "100\n",
		},
		"TeSTusEra (incorrect case)": {
			username: "TeSTusEra",
			response: "100\n",
		},
		"Doesn't exist": {
			username: "Doesn't exist",
			err:      fmt.Errorf("Not found"),
		},
		"Empty": {
			response: "TestUserA - 100\nTestUserB - 100\n",
		},
	}
	for name, tc := range testcases {
		resp, e := assembly.GetBankBalance(tc.username, testUsers)
		if tc.err != nil {
			assert.Errorf(t, tc.err, e.Error(), "%s returned and incorrect error", name)
		}
		assert.Equalf(t, tc.response, resp, "%s response was incorrect", name)
	}

}

func TestBid(t *testing.T) {

	testcases := map[string]struct {
		seller      *assembly.User
		buyer       *assembly.User
		owner       *assembly.User
		book        *assembly.Book
		arguments   string
		buyerBal    *money.Money
		sellerBal   *money.Money
		status      bool
		expectedErr error
	}{
		"Successful Bid": {
			seller:    testUsers[0],
			buyer:     testUsers[1],
			owner:     testUsers[1],
			book:      testBooks[0],
			arguments: "TestUserB, TestBook, 10",
			buyerBal:  money.New(9000, "AUD"),
			sellerBal: money.New(11000, "AUD"),
			status:    false,
		},
		"Bid too low": {
			seller:    testUsers[0],
			owner:     testUsers[0],
			buyer:     testUsers[1],
			book:      testBooks[0],
			status:    true,
			buyerBal:  money.New(10000, "AUD"),
			sellerBal: money.New(10000, "AUD"),
			arguments: "TestUserB, TestBook, 7",
		},
		"Insufficient funds for bid": {
			seller:      testUsers[0],
			owner:       testUsers[0],
			buyer:       testUsers[1],
			book:        testBooks[0],
			status:      true,
			buyerBal:    money.New(10000, "AUD"),
			sellerBal:   money.New(10000, "AUD"),
			arguments:   "TestUserB, TestBook, 100000",
			expectedErr: errors.New("Sorry, you do not have enough available funds to make that bid"),
		},
		"Buyer is owner": {
			seller:      testUsers[0],
			owner:       testUsers[0],
			buyer:       testUsers[1],
			book:        testBooks[0],
			status:      true,
			buyerBal:    money.New(10000, "AUD"),
			sellerBal:   money.New(10000, "AUD"),
			arguments:   "TestUserA, TestBook, 7",
			expectedErr: errors.New("TestBook cannot be sold to TestUserA"),
		},
		"Book doesn't exist": {
			seller:      testUsers[0],
			owner:       testUsers[0],
			buyer:       testUsers[1],
			book:        testBooks[0],
			status:      true,
			buyerBal:    money.New(10000, "AUD"),
			sellerBal:   money.New(10000, "AUD"),
			arguments:   "TestUserA, TestBookXYZ is wrong, 7",
			expectedErr: errors.New("No book with that title found"),
		},
		"Missing an argument": {
			arguments:   "TestUserA TestBookXYZ is wrong, 7",
			expectedErr: errors.New("Usage: bid(username, bookname, bidAmount)"),
		},
		"Too many arguments": {
			arguments:   "TestUserA, TestBookXYZ is wrong, Extra, 7",
			expectedErr: errors.New("Usage: bid(username, bookname, bidAmount)"),
		},
	}
	for name, tc := range testcases {
		// reset the users and books
		testReset()
		// run test
		testErr := assembly.Bid(tc.arguments, testBooks, testUsers, &testTransactions)
		if tc.expectedErr != nil {
			assert.Equalf(t, tc.expectedErr.Error(), testErr.Error(), "%s returned an incorrect error message", name)
		} else {
			assert.NoErrorf(t, testErr, "%s should not have an error but produced %v", name, testErr)
			// check balances
			// Seller
			assert.Equalf(t, tc.sellerBal.Amount(), tc.seller.BankBal.Amount(), "%s seller balance was incorrect", name)
			// Buyer
			assert.Equalf(t, tc.buyerBal.Amount(), tc.buyer.BankBal.Amount(), "%s buyer balance was incorrect", name)
			// check book ownership
			assert.Equal(t, tc.owner, tc.book.Owner, "Book ownership wasn't updated for %s", name)
			// check book for sale status
			assert.Equal(t, tc.status, tc.book.ForSale, "Book status wasn't updated for %s", name)
		}
	}
}
