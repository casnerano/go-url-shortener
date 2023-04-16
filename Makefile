go-all-tests: go-test go-vet
go-test:
	go test -coverprofile=cover.out -cover ./... && go tool cover -func=cover.out
go-vet:
	go vet -vettool=${GOPATH}/bin/statictest.exe ./...
docker-compose-up:
	docker-compose -f ./infrastructure/docker-compose.yaml up -d
migrate:
	migrate -database postgres://gofer:gofer@localhost:5432/shortener?sslmode=disable -source file://migrations/postgres up
