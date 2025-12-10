.PHONY: build fmt vet check clean install

# Build binary
build:
	go build -o bin/mali ./cmd/main.go

# Format code
fmt:
	gofmt -s -w .

# Run vet
vet:
	go vet ./...

# Run all checks
check: fmt vet

# Install to GOPATH/bin
install: build
	@mkdir -p $$(go env GOPATH)/bin
	cp bin/mali $$(go env GOPATH)/bin/mali
	@echo "Installed to $$(go env GOPATH)/bin/mali"

# Clean build artifacts
clean:
	go clean
	rm -rf bin
