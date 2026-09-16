# Makefile модуля go-samples.

BIN_DIR := $(CURDIR)/bin
PACKAGES := hello-go hello-rest-api url-shortener

.PHONY: all build hello-go hello-rest-api url-shortener run-% clean generate help

all: build

# Собрать все бинарники в ./bin.
build: $(PACKAGES)

hello-go:
	@$(MAKE) -C hello-go build

hello-rest-api:
	@$(MAKE) -C hello-rest-api build

url-shortener:
	@$(MAKE) -C url-shortener build

# Примеры: make run-hello-go, make run-hello-rest-api.
run-%:
	@$(MAKE) -C $* run

generate:
	@$(MAKE) -C hello-rest-api generate
	@$(MAKE) -C url-shortener generate

clean:
	@$(MAKE) -C hello-go clean
	@$(MAKE) -C hello-rest-api clean
	@$(MAKE) -C url-shortener clean

help:
	@echo "Доступные цели:"
	@echo "  make build              — собрать все бинарники в ./bin"
	@echo "  make hello-go           — собрать hello-go"
	@echo "  make hello-rest-api     — собрать REST API"
	@echo "  make url-shortener      — собрать URL shortener"
	@echo "  make run-hello-go       — запустить hello-go"
	@echo "  make run-hello-rest-api — запустить REST API"
	@echo "  make run-url-shortener  — запустить URL shortener"
	@echo "  make generate           — обновить ogen-код"
	@echo "  make clean              — удалить бинарники"
