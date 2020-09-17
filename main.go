package main

import (
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
	var bytes []byte
	var err error
	if filename == "" || filename == "-" {
		bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, err = ioutil.ReadFile(filename)
	}
	if err != nil {
		logger.Fatalf("Failed to read file: %s\n", err)
	}

	content, err := convert.Bytes(bytes, filename, options)
	if err != nil {
		logger.Fatalf("Failed to convert file: %v", err)
	}

	os.Stdout.Write(content)
}
