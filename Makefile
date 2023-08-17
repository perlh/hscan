# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOGET = $(GOCMD) get
BINARY_NAME = hscan
BIN_DIR = build
# 编译参数
ARGS = -ldflags="-s -w" -trimpath
# Main build target
all: clean build

# Build target
build: build-linux-arm64 build-linux-amd64 build-linux-mips64 build-windows-amd64 build-windows-arm64 build-windows-386 build-darwin-arm64 build-darwin-amd64

build-linux-arm64:
	$(eval GOOS=linux)
	$(eval ARCH=arm64)
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) $(ARGS) -o $(BIN_DIR)/$(BINARY_NAME)_$(GOOS)_$(ARCH)

build-linux-amd64:
	$(eval GOOS=linux)
	$(eval ARCH=amd64)
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) $(ARGS) -o $(BIN_DIR)/$(BINARY_NAME)_$(GOOS)_$(ARCH)

build-linux-mips64:
	$(eval GOOS=linux)
	$(eval ARCH=mips64)
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) $(ARGS) -o $(BIN_DIR)/$(BINARY_NAME)_$(GOOS)_$(ARCH)

build-windows-amd64:
	$(eval GOOS=windows)
	$(eval ARCH=amd64)
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) $(ARGS) -o $(BIN_DIR)/$(BINARY_NAME)_$(GOOS)_$(ARCH).exe

build-windows-386:
	$(eval GOOS=windows)
	$(eval ARCH=386)
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) $(ARGS) -o $(BIN_DIR)/$(BINARY_NAME)_$(GOOS)_$(ARCH).exe

build-windows-arm64:
	$(eval GOOS=windows)
	$(eval ARCH=arm64)
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) $(ARGS) -o $(BIN_DIR)/$(BINARY_NAME)_$(GOOS)_$(ARCH).exe

build-darwin-arm64:
	$(eval GOOS=darwin)
	$(eval ARCH=arm64)
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) $(ARGS) -o $(BIN_DIR)/$(BINARY_NAME)_$(GOOS)_$(ARCH)

build-darwin-amd64:
	$(eval GOOS=darwin)
	$(eval ARCH=amd64)
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) $(ARGS) -o $(BIN_DIR)/$(BINARY_NAME)_$(GOOS)_$(ARCH)

#build-mac:
#	@echo "Building $(BINARY_NAME) for macOS"
#	@mkdir -p $(BIN_DIR)
#	CGO_ENABLED=0 GOOS=darwin GOARCH=$(ARCH) $(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME)_mac

# Clean target
clean:
	@echo "Cleaning"
	$(GOCLEAN)
	rm -rf $(BIN_DIR)

# Install dependencies
deps:
	$(GOGET) ./...