package assembly_test

import (
	"fmt"
	"testing"

	assembly "github.com/shanehowearth/assemblyPayments/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {
	testcases := map[string]struct {
		testbook *assembly.Book
		first    *assembly.User
		second   *assembly.User
		amount1  string
		amount2  string
		tType1   string
		tType2   string
		expected bool
	}{
		"Successful bid": {
			testbook: testBooks[0],
			first:    testUsers[0],
			second:   testUsers[1],
			amount1:  "10.00",
			amount2:  "TestBook",
			tType1:   "paid",
			tType2:   "sold",
			expected: true,
		},
		"Unsuccessful bid": {
			testbook: testBooks[0],
			first:    testUsers[0],
			second:   testUsers[1],
			amount1:  "1.00",
			expected: false,
		},
	}
	for name, tc := range testcases {
		testReset()
		// create a tranasaction
		arguments := fmt.Sprintf("%s, %s, %s", tc.second.Name, tc.testbook.Title, tc.amount1)
		_ = assembly.Bid(arguments, testBooks, testUsers, &testTransactions)
		// GetTransactions returns a string.
		// However the tests will check the transactions objects directly
		_ = assembly.GetTransactions(&testTransactions)
		fmt.Println(testTransactions)
		if tc.expected {
			assert.NotEmptyf(t, testTransactions, "%s has produced no transactions", name)
			assert.Equal(t, tc.second.Name, testTransactions[0].First.Name, "%s has the wrong user first in the first entry", name)
			assert.Equal(t, tc.first.Name, testTransactions[0].Second.Name, "%s has the wrong user second in the first entry", name)
			assert.Equal(t, tc.amount1, testTransactions[0].Value, "%s has the wrong value in the first entry", name)
			assert.Equal(t, tc.tType1, testTransactions[0].TType, "%s has the wrong transaction type in the first entry", name)
			assert.Equal(t, tc.first.Name, testTransactions[1].First.Name, "%s has the wrong user first in the second entry", name)
			assert.Equal(t, tc.second.Name, testTransactions[1].Second.Name, "%s has the wrong user second in the second entry", name)
			assert.Equal(t, tc.amount2, testTransactions[1].Value, "%s has the wrong value in the second entry", name)
			assert.Equal(t, tc.tType2, testTransactions[1].TType, "%s has the wrong transaction type in the second entry", name)
		} else {
			assert.Emptyf(t, testTransactions, "%s has produced transactions when none were expected", name)

		}
	}
}
