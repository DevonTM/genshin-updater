# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
BINARY_NAME = genshin-updater

# Check the operating system and set the binary name accordingly
ifeq ($(OS),Windows_NT)
	BINARY_NAME := $(BINARY_NAME).exe
endif

.PHONY: all build clean test

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v -ldflags "-s -w" ./main.go

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)
