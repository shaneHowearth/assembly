all:
	go build -o assemblyPayments cmd/main.go
	docker build -t "assemblypayments:shane" .
	docker run -t -d --name=shane assemblypayments:shane

