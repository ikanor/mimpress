package main

import "bytes"

var slideTemplate []byte = []byte("<div class=\"slide\">\n{{}}\n</div>")
var slidePlaceholder []byte = []byte("{{}}")

func HTML2Impress(input [][]byte) [][]byte {
	output := make([][]byte, len(input))

	for i, slide := range input {
		output[i] = bytes.Replace(slideTemplate, slidePlaceholder, slide, 1)
	}

	return output
}

func InsertSlides(slides [][]byte, template []byte) []byte {
	return []byte("")
}
