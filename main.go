package main

import (
	"encoding/json"
	"flag"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcl/hclsyntax"
	"github.com/hashicorp/hcl2/hclpack"
	"io/ioutil"
	"log"
	"os"
)

var usePackFormat bool

func init() {
	flag.BoolVar(&usePackFormat, "pack", false, "If set, use the pack json format, instead of valid hcl json")
}

func main() {
	logger := log.New(os.Stderr, "", 0)
	flag.Parse()
	var err error

	var filename = flag.Arg(0)
	var bytes []byte
	if filename == "" || filename == "-" {
		bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, err = ioutil.ReadFile(filename)
	}

	if err != nil {
		logger.Fatalf("Failed to read file: %s\n", err)
	}

	var content interface{}
	if usePackFormat {
		content, err = getPackJSON(bytes, filename)
	} else {
		content, err = getHclJSON(bytes, filename)
	}

	if err != nil {
		logger.Fatalf("Failed to convert file: %v", err)
	}

	jb, err := json.MarshalIndent(content, "", "    ")

	if err != nil {
		logger.Fatalf("Failed to generate JSON: %v", err)
	}

	os.Stdout.Write(jb)

	// ioutil.WriteFile(f+".json", jb, 0666)
}

func getPackJSON(bytes []byte, filename string) (interface{}, error) {
	f, diags := hclpack.PackNativeFile(bytes, filename, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, diags
	}
	return f, nil
}

func getHclJSON(bytes []byte, filename string) (interface{}, error) {
	file, diags := hclsyntax.ParseConfig(bytes, filename, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, diags
	}
	body := file.Body.(*hclsyntax.Body)

	c := converter{bytes: bytes}
	return c.convertBody(body)
}
