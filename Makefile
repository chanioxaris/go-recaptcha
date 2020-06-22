COVERAGE_FILE := "./cover.out"

# Run linter
lint:
	golangci-lint run

# Run vet
vet:
	go vet ./...

# Run all tests
test:
	go test ./...

# Run all tests and get coverage report
coverage:
	go test ./... -coverprofile $(COVERAGE_FILE) && \
	rm $(COVERAGE_FILE)

# Run all tests and get HTML coverage report
coverage-html:
	go test ./... -coverprofile $(COVERAGE_FILE) && \
	go tool cover -html=$(COVERAGE_FILE) && \
	rm $(COVERAGE_FILE)