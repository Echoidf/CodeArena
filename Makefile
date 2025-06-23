build:
	go build -o ./bin/server

run: build
	./bin/server

test:
	go test -v ./...
