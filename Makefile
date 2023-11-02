all: mocks build test

install-dep:
	go install github.com/vektra/mockery/v2@v2.36.0

mocks:
	mockery --all --case snake

build: 
	go build -o ./build/

test:
	go generate
	go test ./... 

install:
	go install