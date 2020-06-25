dev:
	go run httpd/main.go

test: 
	go test -coverprofile fmtcoverage.html fmt
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out

pro:
	go build -o gin httpd/main.go
	nohup ./gin & disown

stop_pro:
	kill -9 $(pgrep -f gin)