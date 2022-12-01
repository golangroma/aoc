build:
	go build -o aoc ./cmd/aoc/main.go

test:
	go test -v $(shell go list ./... | grep -v template)