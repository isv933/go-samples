# Makefile пакета hello-go

BINARY = hello-go-bin
SRC = .

.PHONY: build run clean

# Собрать бинарник
build:
	@echo "Сборка $(BINARY)..."
	go build -o $(BINARY) $(SRC)

# Запустить бинарник
run: build
	@echo "Запуск $(BINARY)..."
	./$(BINARY)

# Удалить бинарник
clean:
	@echo "Удаление $(BINARY)..."
	rm -f $(BINARY)

