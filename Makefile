.PHONY: test build clean


build:
	@go build -o mali ./cmd/main.go

clean:
	@rm -f mali mali-* coverage.out coverage.html

fmt:
	@gofmt -s -w .

vet:
	@go vet ./...

check: fmt vet test

