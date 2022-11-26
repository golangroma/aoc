test:
	go test -v $(shell go list ./... | grep -v template)