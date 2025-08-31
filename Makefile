all: build

build: 
	go build -o bin/dsy cmd/devsyringe/main.go

release: 
	sudo go build -o /usr/bin/dsy cmd/devsyringe/main.go

