test:
	go test --race ./...

install: test
	go install github.com/mdwhatcott/mdopen/cmd/mdopen
