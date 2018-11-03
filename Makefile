.PHONY: test
test:
	dep ensure
	go test -v ./...