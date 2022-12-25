tests: go-test go-vet
go-test:
	@go test -count=1 -cover ./...
go-vet:
	@go vet -vettool=${GOPATH}/bin/statictest.exe ./...
docker-compose-up:
	@docker-compose -f ./docker/docker-compose.yaml up -d
migrate:
	@migrate -database postgres://gofer:gofer@localhost:5432/shortener?sslmode=disable -source file://migrations/postgres up
