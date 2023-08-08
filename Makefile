GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
BINARY_NAME = genshin-updater.exe

.PHONY: all build run dist clean

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

dist:
	mkdir -p release/genshin-patch
	$(GOBUILD) -o release/$(BINARY_NAME) -trimpath -ldflags "-s -w"
	cp aria2.conf release/
	cp LICENSE release/

clean:
	$(GOCLEAN)
	rm -rf release
