# mdopen

Allows to view markdown files in the default browser. For more details, see the API [documentation](https://godoc.org/github.com/mdwhatcott/mdopen).

## CLI

Install:

```bash
go get gopkg.in/romanyx/mdopen.v1/cmd/mdopen
```

Create a markdown file:

```bash
echo "# Hello from markdown" > hello.md
```

View it in the default browser as html:

```bash
mdopen hello.md
```

## Library

Install:

```bash
go get github.com/mdwhatcott/mdopen
```

``` go
package main

import "github.com/mdwhatcott/mdopen"

func main() {
    f := strings.NewReader("# Hello from markdown")

    opener := mdopen.New()
    if err := opener.Open(f); err != nil {
        log.Fatal(err)
    }
}
```
