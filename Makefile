tests:
	go test -count=1 -cover ./...
run-server:
	go run -config=./configs/application.yaml
load-migrations:
	migrate -database postgres://gofer:123456@localhost:5432/shortener?sslmode=disable -source file://migrations up
