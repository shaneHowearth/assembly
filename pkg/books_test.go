package assembly_test

import (
	"testing"

	assembly "github.com/shanehowearth/assemblyPayments/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGetBookList(t *testing.T) {
	testcases := map[string]struct {
		username string
		response string
	}{
		"Successful search": {
			username: "TestUserA",
			response: "TestBook\n",
		},
		"Successful search with Not for sale book": {
			username: "TestUserb",
			response: "TestBookB\nNot For Sale\n",
		},
		"No username provided": {
			response: "TestBook\nTestBookB\n",
		},
		"Random case": {
			username: "TeSTuSeRA",
			response: "TestBook\n",
		},
		"Non Existant": {
			username: "Non-existant",
			response: "User not found",
		},
		"Leading and Trailing whitespace": {
			username: " TestUserA ",
			response: "TestBook\n",
		},
		"Unicode": {
			username: "Testing «ταБЬℓσ»: 1<2 & 4+1>3, now 20% off!",
			response: "User not found",
		},
	}
	testReset()
	for name, tc := range testcases {
		bookString := assembly.GetBookList(tc.username, testBooks, testUsers)
		assert.Equal(t, tc.response, bookString, "%s provided an incorrect book string", name)
	}
}
