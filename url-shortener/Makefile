# Makefile приложения url-shortener.

ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/..)
BIN_DIR := $(ROOT_DIR)/bin
BINARY := $(BIN_DIR)/url-shortener

.PHONY: build run generate clean

build:
	@mkdir -p "$(BIN_DIR)"
	@echo "Сборка $(BINARY)..."
	go build -o "$(BINARY)" ./cmd/server

run: build
	@echo "Запуск $(BINARY)..."
	"$(BINARY)"

generate:
	go generate ./gen/...

clean:
	@echo "Удаление $(BINARY)..."
	rm -f "$(BINARY)"
