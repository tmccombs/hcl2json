package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/tmccombs/hcl2json/convert"
)

func main() {
	logger := log.New(os.Stderr, "", 0)

	var options convert.Options

	flag.BoolVar(&options.Simplify, "simplify", false, "If true attempt to simply expressions which don't contain any variables or unknown functions")
	flag.Parse()

	var filename = flag.Arg(0)
	var fileBytes []byte
	var err error
	if filename == "" || filename == "-" {
		fileBytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		fileBytes, err = ioutil.ReadFile(filename)
	}
	if err != nil {
		logger.Fatalf("Failed to read file: %s\n", err)
	}

	converted, err := convert.Bytes(fileBytes, filename, options)
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
