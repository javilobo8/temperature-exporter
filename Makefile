# build for golang

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build

# Binary name
BINARY_NAME=temperature-exporter

# Build
build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v

# Clean
clean:
	rm -f bin/$(BINARY_NAME)

# Run
run:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v ./...
	./bin/$(BINARY_NAME)