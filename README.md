# Assembly
Small book sales management application

# Installation
Build with `make`

That command will create a container named `shane` that you can access with `docker exec -it shane sh`.

# Usage
Inside the container execute `./assemblyPayments`
Available commands:
+   Help - Prints this message.
+   Bal - Gets the current bank balances.
       If a name is given then only the bank balance for that person.
+   Bid - Bid for a book.
       Give the name of the person bidding, the name of the book they are bidding for, and the amount that they are bidding. If successful the book will change ownership, and the money paid from the bidder's wallet to the seller's.
+   Books - List all the books for sale, or all of the books owned by the given name.
+   Transactions - Get all of the transactions that have occurred.
+   Exit - Exits this program.
