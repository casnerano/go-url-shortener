tests:
	go test -count=1 -cover ./...
docker-compose-up:
	docker-compose -f ./docker/docker-compose.yaml up -d
migrate:
	migrate -database postgres://gofer:gofer@localhost:5432/shortener?sslmode=disable -source file://migrations/postgres up
