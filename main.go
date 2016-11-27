package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Template string `short:"t" long:"template" required:"true"`
	Files    struct {
		Input  string `positional-arg-name:"input file" required:"true"`
		Output string `positional-arg-name:"output file"`
	} `positional-args:"true"`
}

func parseOpts() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
	if opts.Files.Output == "" {
		opts.Files.Output = opts.Files.Input + ".html"
	}
}

func readFile(filename string) []byte {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return contents
}

func writeFile(filename string, contents []byte) {
}

func main() {
	parseOpts()

	source := readFile(opts.Files.Input)
	template := readFile(opts.Template)

	html := Markdown2HTML(source)
	slides := HTML2Impress(html)
	outcome := InsertSlides(slides, template)

	writeFile(opts.Files.Output, outcome)
}
