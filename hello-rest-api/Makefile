# Makefile приложения hello-rest-api.

ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/..)
BIN_DIR := $(ROOT_DIR)/bin
BINARY := $(BIN_DIR)/hello-rest-api

.PHONY: build run generate clean

build:
	@mkdir -p "$(BIN_DIR)"
	@echo "Сборка $(BINARY)..."
	go build -o "$(BINARY)" ./cmd/server

run: build
	@echo "Запуск $(BINARY)..."
	"$(BINARY)"

generate:
	go generate ./gen/api

clean:
	@echo "Удаление $(BINARY)..."
	rm -f "$(BINARY)"
