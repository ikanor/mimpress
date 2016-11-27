package main

import (
	"bytes"

	"github.com/russross/blackfriday"
)

var SLIDE_SEPARATOR string = "***"

func Markdown2HTML(input []byte) [][]byte {
	slides := bytes.Split(input, []byte(SLIDE_SEPARATOR))
	output := make([][]byte, len(slides))

	for i, slide := range slides {
		output[i] = blackfriday.MarkdownCommon(slide)
	}

	return output
}
