GO ?= go
CP := cp
RM := rm -rf
BINARY_NAME := genshin-updater
OUT_DIR := release

ifeq ($(OS),Windows_NT)
	CP := xcopy /Q /Y
	RM := cmd /C RD /Q /S
endif

ifeq ($(shell $(GO) env GOOS),windows)
	BINARY_NAME := $(BINARY_NAME).exe
endif

.PHONY: all build clean dist run

all: dist

build:
	$(GO) build -v -trimpath -ldflags "-s -w" -o $(OUT_DIR)/$(BINARY_NAME)

run: build
	$(OUT_DIR)/$(BINARY_NAME)

dist: build
	$(CP) aria2.conf $(OUT_DIR)
	$(CP) LICENSE $(OUT_DIR)

clean:
	$(RM) $(OUT_DIR)
