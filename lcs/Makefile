.PHONY: build test clean

build: deps
	go build -v

test: build
	go test -v

deps:
	go get -v

clean:
	go clean
