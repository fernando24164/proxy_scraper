.PHONY: test
test:
	go test -v ./...

.PHONY: deps
deps:
	go get ./...