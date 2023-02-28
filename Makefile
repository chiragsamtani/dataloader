test:
	go test ./...
build-main:
	go build -o bin/main main.go
docs:
	godoc -goroot=${PWD}/internal/
run:
	go run main.go
docker-clean:
	docker-compose stop
	docker-compose down -v
docker:
	docker-compose down
	docker-compose up --build -d
