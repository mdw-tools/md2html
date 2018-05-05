package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mdwhatcott/mdopen"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: mdopen <files>")
		flag.PrintDefaults()
	}
	flag.Parse()
	opener := mdopen.New(mdopen.GithubTemplate())
	for _, path := range flag.Args() {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err, "failed to open file")
		}

		if err := opener.Open(file); err != nil {
			log.Fatal(err, "open markdown failed")
		}
	}
}
