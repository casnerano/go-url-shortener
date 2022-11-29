tests:
	go test -count=1 -cover ./...
run-server:
	go run -config=./configs/application.yaml
