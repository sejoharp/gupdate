ifndef VERBOSE
.SILENT:
endif

binary="gupdate"

.PHONY: run build

run: debug-build ## Build and run binary without arguments
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

debug-build: ## Build binary
	go build -o $(binary)

test: ## Run tests
	go test ./...

dependencies:
	go get ./...

cover: ## Run test-coverage and open in browser
	go test -v -covermode=count -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

quick-cover: ## Run simple coverage
	go test -cover ./...

fmt: ## Format source-tree
	gofmt -l -s -w .

help: ## Print all available make-commands
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
