package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"

	"github.com/tmccombs/hcl2json/convert"
)

func main() {
	logger := log.New(os.Stderr, "", 0)

	var options convert.Options

	flag.BoolVar(&options.Simplify, "simplify", false, "If true attempt to simply expressions which don't contain any variables or unknown functions")
	flag.Parse()

	buffer := bytes.NewBuffer([]byte{})
	files := flag.Args()
	var inputName string

	switch len(files) {
	case 0:
		files = append(files, "-")
		inputName = "STDIN"
	case 1:
		inputName = files[0]
		if inputName == "-" {
			inputName = "STDIN"
		}
	default:
		inputName = "COMPOSITE"
	}

	for _, filename := range files {
		var stream io.Reader
		if filename == "-" {
			stream = os.Stdin
			filename = "STDIN" // for better error message
		} else {
			file, err := os.Open(filename)
			if err != nil {
				logger.Fatalf("Failed to open %s: %s\n", filename, err)
			}
			defer file.Close()
			stream = file
		}
		_, err := buffer.ReadFrom(stream)
		if err != nil {
			logger.Fatalf("Failed to read from %s: %s\n", filename, err)
		}
		buffer.WriteByte('\n') // just in case it doesn't have an ending newline
	}

	converted, err := convert.Bytes(buffer.Bytes(), inputName, options)
	if err != nil {
		logger.Fatalf("Failed to convert file: %v", err)
	}

	var indented bytes.Buffer
	if err := json.Indent(&indented, converted, "", "    "); err != nil {
		logger.Fatalf("Failed to indent file: %v", err)
	}

	if _, err := indented.WriteTo(os.Stdout); err != nil {
		logger.Fatalf("Failed to write to standard out: %v", err)
	}
}
