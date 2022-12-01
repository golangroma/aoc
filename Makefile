build:
	cd cli && go build -o ../aoc .

test:
	cd cli && go test -v $(shell cd cli && go list ./... | grep -v template)

tidy:
	find . -name go.mod -execdir go mod tidy \;
