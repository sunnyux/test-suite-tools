# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
MAIN_DIR=test-suite-tools/cmd/tst
DEPS=github.com/urfave/cli

all: build

build:
	$(GOBUILD) $(MAIN_DIR)
install:
	$(GOINSTALL) $(MAIN_DIR)
test:
	$(GOTEST) -v ./...

get:
	$(GOGET) $(DEPS)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
