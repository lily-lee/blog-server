.PHONY: build test run

build:
	go build .

dir ?=./...
test:
	@echo $(ENV)
	go test $(dir) -v -gcflags=-l -p 1 -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

arg ?=blog
run:
	@echo $(ENV)
	go run main.go $(arg)