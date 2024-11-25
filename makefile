run:
	go run cmd/main.go
update-lib:
	go get -u ./... && go mod tidy
test-all:
	go test ./... -cover