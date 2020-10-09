BIN_PATH=bin
BIN_NAME=crate

test:
	go test -v ./... 

build:
	@go build -o $(BIN_PATH)/$(BIN_NAME) main.go

run: build
	@./$(BIN_PATH)/$(BIN_NAME)

dev: 
	modd
	
