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

	content, err := convert.Bytes(bytes, filename)
	if err != nil {
		logger.Fatalf("Failed to convert file: %v", err)
	}

	os.Stdout.Write(content)
}
