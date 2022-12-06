# «Сервис сокращения URL»
Я.Практикум: Курс — Golang advanced

### Настройки приложения

- `server_addr` адрес на котором будет запущен сервер (по умолчанию `localhost:8080`)
- `storage`
  - `type` тип хранилища [`memory`, `database`] (по умолчанию `memory`)
  - `dsn` адрес базы данных
- `short_url`
  - `ttl` время жизни хранения ссылок (по умолчанию `0` — значит храним всегда)

Приложение можно запускать со своим файлом конфигурации в формате json или yaml.

[Примеры](./configs) файлов конфигурации.

Пример запуска приложения со своими настройками:
```bash
cp ./configs/application.yaml.dist ./configs/application.yaml
go run cmd/shortener/main.go -config=./configs/application.yaml

```

### Хранилище ссылок

Если приложение использует Базу данных,
то необходимо заранее создать базу данных и запустить миграции.

Для тестирования, можно запустить Базу данных из докера:
```bash
make docker-compose-up
make migrate
```

или

```bash
docker-compose -f ./docker/docker-compose.yaml up -d
migrate -database postgres://gofer:gofer@localhost:5432/shortener?sslmode=disable -source file://migrations/postgres up
```
