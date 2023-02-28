test:
	go test ./...
build-main:
	go build -o bin/main main.go
docs:
	godoc -goroot=${PWD}/internal/
run:
	go run main.go
