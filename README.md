# go-samples

Небольшой набор учебных примеров на Go. Все приложения находятся в одном модуле и собираются через корневой `Makefile`.

## Примеры

| Каталог | Описание |
| --- | --- |
| `hello-go` | Консольный пример работы с методом и `slog`. |
| `hello-rest-api` | HTTP API для получения пользователя. Обработчик возвращает демонстрационные данные; серверный код сгенерирован из OpenAPI с помощью `ogen`. |
| `url-shortener` | HTTP API для создания и получения коротких URL с хранением в PostgreSQL. Код API также сгенерирован из OpenAPI. |

## Требования

- Go 1.27.1 (версия указана в `go.mod`);
- `make` для команд ниже;
- PostgreSQL для `url-shortener`.

## Сборка и запуск

Из корня репозитория:

```sh
make build                 # собрать все приложения в bin/
make run-hello-go         # вывести сообщение в консоль
make run-hello-rest-api   # запустить HTTP API пользователей
make run-url-shortener    # запустить сервис коротких URL
```

Можно собрать отдельный пример командой `make hello-go`, `make hello-rest-api` или `make url-shortener`. `make clean` удаляет собранные бинарники.

### REST API пользователей

По умолчанию сервер слушает порт `18080`; адрес можно изменить переменной `HTTP_ADDR`:

```sh
HTTP_ADDR=:8080 make run-hello-rest-api
curl http://localhost:8080/users/1
```

Для ID `1` возвращается пользователь Alice, для других ID — демонстрационное имя вида `User 2`.

### Сервис коротких URL

Перед запуском создайте в PostgreSQL таблицу, которую ожидает репозиторий:

```sql
CREATE TABLE shortener_url (
    id text PRIMARY KEY,
    url text NOT NULL
);
```

Скопируйте `url-shortener/gen/cmd/server/settings/url-shortener.sample.json` в свой файл настроек и задайте `listen_address`, `db_connection_string` и `short_link_url` для своего окружения. Значение подключения в образце указывает на `server.lan` и не рассчитано на произвольную локальную установку. Параметр `server_timeout` задается в формате ISO 8601, например `PT5S`.

```sh
make url-shortener
./bin/url-shortener -config-file ./path/to/settings.json
```

Если запустить через `make run-url-shortener` без файла настроек, используются значения по умолчанию из `url-shortener/cmd/server/settings/settings.go`.

API по умолчанию слушает порт `18000`:

```sh
curl -X PUT 'http://localhost:18000/shortener?url=https%3A%2F%2Fexample.com'
curl 'http://localhost:18000/shortener/redirect/ID'
```

Подставьте полученный ID во второй запрос. `GET /shortener/redirect/{id}` возвращает исходный URL как текст, но не выполняет HTTP-перенаправление. `DELETE` для этого пути пока не удаляет запись. Формирование ссылки в текущей реализации также требует доработки, поэтому при проверке используйте ID из ответа на `PUT`.

## Генерация кода и проверка

Спецификации API находятся в `hello-rest-api/api/openapi.yaml` и `url-shortener/api/shortener.yaml`. После их изменения можно обновить сгенерированный код командой `make generate` (она запускает `go generate` и загружает `ogen@latest`).

```sh
go test ./...
```
