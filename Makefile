BIN_NAME=pynezz-cli
#TEST_BIN_NAME=testmodule

WINDOWS=$(BIN_NAME)_win_amd64.exe
LINUX=$(BIN_NAME)_linux_amd64.out

#TEST_WINDOWS=$(TEST_BIN_NAME)_win_amd64.exe
#TEST_LINUX=$(TEST_BIN_NAME)_linux_amd64.out

VERSION=$(shell git describe --tags --always --long)

.PHONY: all test clean

$(LINUX): main.go
	CGO_ENABLED=1 GOARCH=amd64 GOOS=linux CC="zig cc -target x86_64-linux-gnu.2.31.0" CXX="zig c++ -target x86_64-linux-gnu.2.31.0" go build -v -o $(LINUX) -tags linux -ldflags="-s -w -X main.buildVersion=$(VERSION)" .
# CGO_ENABLED=15.10.218-1
# $(TEST_LINUX): cmd/testmodule/main.go
# 	CGO_ENABLED=1 go build -v -o $(TEST_LINUX) -tags linux -ldflags="-s -w -X main.buildVersion=$(VERSION)" ./cmd/testmodule/main.go

$(WINDOWS): main.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -v -o $(WINDOWS) -ldflags="-s -w -X main.buildVersion=$(VERSION)" .

# $(TEST_WINDOWS): cmd/testmodule/main.go
# 	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -v -o $(TEST_WINDOWS) -ldflags="-s -w -X main.buildVersion=$(VERSION)" ./cmd/testmodule/main.go

# Build targets
windows: $(WINDOWS)
linux: $(LINUX)
# prototype: $(TEST_LINUX) $(LINUX)

test:	## Run tests
	gen
	go test ./...

vet: ## Run go vet
	go vet ./...

build: gen windows linux ## Build the application for Windows and Linux
	@echo $(VERSION)
	@echo "Build complete"

build-windows: windows ## Build the application for Windows
	@echo $(VERSION)
	@echo "Build complete"

build-linux: linux  ## Build the application for Linux
	@templ generate
	@echo $(VERSION)
	@echo "Build complete"

run: ## Build and run the application (Linux)
	$(LINUX) && ./$(LINUX)

gen: ## Generate code
	@templ generate

gen-run: gen ## Generate code and run the application
	go run . serve -p 8080


# run-prototype: # Run prototype
# 	$(TEST_LINUX)
# 	./$(TEST_LINUX)

clean:	## Remove build files
	go clean
	rm $(WINDOWS) $(LINUX)
# $(TEST_WINDOWS) $(TEST_LINUX)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
