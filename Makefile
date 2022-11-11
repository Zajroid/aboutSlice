BINARY_NAME=main.out

run:
	go run main.go

build-r:
	go build -ldflags "-s -w" -o ./bin/${BINARY_NAME} main.go

build:
	go build -o ./bin/${BINARY_NAME} main.go

deps:
	go get ...