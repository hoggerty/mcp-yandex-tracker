BINARY := yandex-tracker-mcp
VERSION ?= $(shell git describe --tags --always --dirty)
LDFLAGS := -s -w -X github.com/hoggerty/mcp-yandex-tracker/cmd.VERSION=$(VERSION)

.PHONY: build
build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) .

.PHONY: run
run:
	go run -ldflags "$(LDFLAGS)" . server

.PHONY: test
test:
	go test ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	gofmt -l .

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: clean
clean:
	rm -f $(BINARY)
