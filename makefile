dev:
	go run httpd/main.go

test: 
	go test -coverprofile fmtcoverage.html fmt
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out