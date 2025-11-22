package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

var Version = "dev"

func main() {
	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	flags.Usage = func() {
		_, _ = fmt.Fprintf(flags.Output(), "Usage of %s:\n", flags.Name())
		_, _ = fmt.Fprintf(flags.Output(), "Given markdown text input via stdin, produce fully rendered html output via stdout.")
		flags.PrintDefaults()
	}
	_ = flags.Parse(os.Args[1:])

	converter := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			extension.GFM,
			extension.DefinitionList,
			extension.Footnote,
		),
	)
	markdown, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	var HTML bytes.Buffer
	err = converter.Convert(markdown, &HTML)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("<html>\n<head></head>\n<body>\n%s\n</body>\n</html>", HTML.String())
}
