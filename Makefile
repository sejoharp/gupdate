ifndef VERBOSE
.SILENT:
endif

.DEFAULT_GOAL:= help
binary="gupdate"

.PHONY: install
install: ## Install dependencies
ifeq ($(CI),true)
	$(info "Running in CI mode")
else
	pre-commit install
	go mod download
	go mod tidy
	go mod verify
endif

.PHONY: pre-commit-install
pre-commit-install: ## install pre-commit
	pre-commit install

.PHONY: pre-commit
pre-commit: pre-commit-install ## run pre-commit
	pre-commit run --all-files

.PHONY: run
run: install debug-build ## Build and run binary without arguments
	./$(binary)

.PHONY: build-linux-amd64
build-linux-amd64: ## Build binary
	env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(binary)-linux-amd64

.PHONY: build-darwin-amd64
build-darwin-amd64: ## Build binary
	env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o $(binary)-darwin-amd64

.PHONY: build-darwin-arm64
build-darwin-arm64: ## Build binary
	env GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o $(binary)-darwin-arm64

.PHONY: build-windows-amd64
build-windows-amd64: ## Build binary
	env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o $(binary)-windows-amd64

.PHONY: install-linux-amd64
install-linux-amd64: build-linux-amd64 ## install to home bin directory
	mv $(binary)-linux-amd64 ~/bin/gupdate

.PHONY: install-darwin-amd64
install-darwin-amd64: build-darwin-amd64 ## install to home bin directory
	mv $(binary)-darwin-amd64 ~/bin/gupdate

.PHONY: install-darwin-arm64
install-darwin-arm64: build-darwin-arm64 ## install to home bin directory
	mv $(binary)-darwin-arm64 ~/bin/gupdate

.PHONY: install-windows-amd64
install-windows-amd64: build-windows-amd64 ## install to home bin directory
	mv $(binary)-windows-amd64 ~/bin/gupdate

.PHONY: debug-build
debug-build: ## Build binary
	go build -o $(binary)

.PHONY: test
test: install ## Run tests
	go test ./...

.PHONY: dependencies
dependencies: ## Install dependencies
	go get ./...

.PHONY: cover
cover: ## Run test-coverage and open in browser
	go test -v -covermode=count -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

.PHONY: quick-cover
quick-cover: ## Run simple coverage
	go test -cover ./...

.PHONY: fmt
fmt: ## Format source-tree
	gofmt -l -s -w .

CURRENT_VERSION=$(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
MAJOR=$(shell echo $(CURRENT_VERSION) | sed 's/v//' | cut -d. -f1)
MINOR=$(shell echo $(CURRENT_VERSION) | sed 's/v//' | cut -d. -f2)
PATCH=$(shell echo $(CURRENT_VERSION) | sed 's/v//' | cut -d. -f3)

NEXT_MAJOR=v$(shell expr $(MAJOR) + 1).0.0
NEXT_MINOR=v$(MAJOR).$(shell expr $(MINOR) + 1).0
NEXT_PATCH=v$(MAJOR).$(MINOR).$(shell expr $(PATCH) + 1)

define release
	@git tag -a $(1) -m "Release $(1)"
	@git push origin $(1)
	@echo "Released $(1)"
endef

.PHONY: major
major:
	$(call release,$(NEXT_MAJOR))

.PHONY: minor
minor:
	$(call release,$(NEXT_MINOR))

.PHONY: patch
patch:
	$(call release,$(NEXT_PATCH))

.PHONY: help
help: ## Print all available make-commands
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
