YEAR ?= 2024

buildall:
	go build -o bin ./$(YEAR)/days/...
.PHONY: buildall

testall:
	go test ./...
.PHONY: testall
