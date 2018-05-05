test:
	go test --race ./...

install: test template
	go install github.com/mdwhatcott/mdopen/cmd/mdopen

template:
	cd $(GOPATH)/src/github.com/mdwhatcott/mdopen/internal/templates/github; go run make_template.go