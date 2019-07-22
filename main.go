package main

import (
	"encoding/json"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hclpack"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "", 0)
	f := os.Args[1]
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		logger.Fatalf("Uh oh: %s\n", err)
	}

	body, diags := hclpack.PackNativeFile(bytes, f, hcl.Pos{Line: 1, Column: 1})

	if diags.HasErrors() {
		logger.Fatalf("Failed to parse: %s", diags.Error())
	}

	jb, err := json.MarshalIndent(body, "", "    ")
	if err != nil {
		logger.Fatalf("Failed to marshal: %v", err)
	}

	os.Stdout.Write(jb)

	// ioutil.WriteFile(f+".json", jb, 0666)
}
