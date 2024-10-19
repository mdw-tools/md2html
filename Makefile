test:
	go test --race ./...

install: test
	go install github.com/mdw-tools/mdopen/cmd/mdopen
