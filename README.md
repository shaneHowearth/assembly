# Assembly
Small book sales management application

# Installation
Run with `go run assemblyPayments.go books.go transactions.go user.go`
or
Build with `go build assemblyPayments.go books.go transactions.go user.go`


# Usage
Available commands:
+   Help - Prints this message.
+   Bal - Gets the current bank balances.
       If a name is given then only the bank balance for that person.
+   Bid - Bid for a book.
       Give the name of the person bidding, the name of the book they are bidding for, and the amount that they are bidding. If successful the book will change ownership, and the money paid from the bidder's wallet to the seller's.
+   Books - List all the books for sale, or all of the books owned by the given name.
+   Exit - Exits this program.
