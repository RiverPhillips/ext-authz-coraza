export GOBIN ?= $(shell pwd)/bin

GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

GOLANGCI = $(GOBIN)/golangci-lint

.PHONY: build
build: install
	go build ./...

.PHONY: install
install:
	go mod download

.PHONY: test
test:
	go test -race ./...

.PHONY: cover
cover:
	go test -coverprofile=cover.out -covermode=atomic -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html

$(GOLANGCI): tools/go.sum
	cd tools && go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint: $(GOLANGCI)
	@rm -rf lint.log
	@echo "Checking gofmt"
	@gofmt -d -s $(GO_FILES) 2>&1 | tee lint.log
	@echo "Checking go vet"
	@go vet ./... 2>&1 | tee -a lint.log
	@echo "Checking golangci-lint"
	@$(GOLANGCI) run ./... | tee -a lint.log
	@[ ! -s lint.log ]
