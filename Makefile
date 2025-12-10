<<<<<<< Updated upstream
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


=======
.PHONY: build fmt vet check clean

build:
	go build -o bin/mali ./cmd/main.go

fmt:
	gofmt -s -w .

vet:
	go vet ./...

check: fmt vet

clean:
	go clean
	rm -rf bin
>>>>>>> Stashed changes
