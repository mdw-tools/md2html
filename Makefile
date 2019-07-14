test:
	go test --race ./...

install: test template
	go install github.com/mdwhatcott/mdopen/cmd/mdopen

template:
	cd internal/templates/github; go run make_template.go
