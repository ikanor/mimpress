package main

import (
	"bytes"
	"text/template"
)

func InsertSlides(slides [][]byte, tmplt []byte) ([]byte, error) {
	t, err := template.New("presentation").Parse(string(tmplt))
	if err != nil {
		return nil, err
	}

	var output bytes.Buffer
	params := struct{ Slides []string }{byteArrToStringArr(slides)}
	err = t.Execute(&output, params)
	if err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}

func byteArrToStringArr(input [][]byte) []string {
	strings := make([]string, len(input))
	for i, bytes := range input {
		strings[i] = string(bytes)
	}
	return strings
}
