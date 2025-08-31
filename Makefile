all: build

build: 
	go build -o bin/dsy cmd/dsy/main.go

release: 
	sudo go build -o /usr/bin/dsy cmd/dsy/main.go

