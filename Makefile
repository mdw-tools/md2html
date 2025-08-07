test:
	go test --race ./...

install: test
	go install github.com/mdw-tools/md2html/cmd/md2html
