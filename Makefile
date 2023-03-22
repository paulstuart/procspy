.PHONY: all build test install bench
all: test build install

build: 
	go vet
	#golangci-lint run ./...
	cd cmd/procspy; go build

# buildall:
# 	GOOS=darwin go build
# 	GOOS=linux go build

test:
	go test

install:
	cd cmd/procspy; go install

bench:
	go test -bench .
