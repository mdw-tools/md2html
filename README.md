# md2html

Convert markdown to html (for viewing in browser, printing to pdf, etc.).

## Rudimentary usage

```shell
$ cat content.md | md2html > content.html
```

## More advanced usage

Requires [`entr`](https://github.com/eradman/entr)

```shell
$ ls content.md | entr -cs 'cat content.md | md2html > content.html'
```

Then navigate to file:///path/to/content.html in a browser. Whenever the file is saved, the html will be regenerated, and refreshing the browser will load the new content. (FUTURE: make the browser auto-reload, maybe via a websockets connection.) 