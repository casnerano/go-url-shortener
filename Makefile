tests:
	go test -count=1 -cover ./...
migrate:
	migrate -database postgres://gofer:gofer@localhost:5432/shortener?sslmode=disable -source file://migrations up
