# Корневой Makefile модуля

# Список папок с пакетами, которые нужно собирать.
# Добавляйте сюда новые пакеты по мере появления.
PACKAGES = hello-go

.PHONY: all build clean run help $(PACKAGES)

# Цель по умолчанию — собрать все пакеты
all: build

# Собрать все пакеты
build: $(PACKAGES)

# Запустить make в каждой папке пакета
$(PACKAGES):
	@echo "==> Building $@"
	@$(MAKE) -C $@ build

# Запустить конкретный пакет (например: make run-hello-go)
run-%:
	@$(MAKE) -C $* run

# Очистить все пакеты
clean:
	@for pkg in $(PACKAGES); do \
		echo "==> Cleaning $$pkg"; \
		$(MAKE) -C $$pkg clean; \
	done

# Показать список доступных целей
help:
	@echo "Доступные цели:"
	@echo "  make build          — собрать все пакеты"
	@echo "  make clean          — очистить все пакеты"
	@echo "  make run-hello-go   — запустить пакет hello-go"
	@echo "  make help           — показать эту справку"

