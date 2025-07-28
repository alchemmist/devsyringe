all: build

build: 
	go build -o bin/dsy cmd/devsyringe/main.go

