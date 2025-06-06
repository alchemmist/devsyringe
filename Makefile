all: build

build: 
	go build -o bin/devsyringe cmd/devsyringe/main.go
