# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=dirwalk
BINARY_LINUX=$(BINARY_NAME)_linux

all: test tidy build
build: 
		CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) -v
tidy:
		$(GOMOD) tidy
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_LINUX)
run: build
		./$(BINARY_NAME)

# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX) -v
