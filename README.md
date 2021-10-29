# hcl2json

This is a tool to convert from [HCL](https://github.com/hashicorp/hcl2/blob/master/hcl/hclsyntax/spec.md) to json, to make it
easier for non-go applications and scripts to process HCL inputs (such as terraform config).

If passed the `-pack` option, it converts to the JSON pack format used by
[hclpack](https://godoc.org/github.com/hashicorp/hcl2/hclpack), which contains the original structure of the document. However,
all it has for expressions is the source text, which isn't always very useful.

If no options are passed, it converts the native HCL file to an (almost) equivalent HCL JSON file. Note, however, that there are
some corner cases where it may not be exactly equivalent, especially if the target application makes use of [static
analysis](https://github.com/hashicorp/hcl2/blob/master/hcl/hclsyntax/spec.md#static-analysis).

## Alternatives

At the time this project created, there weren't really any other object.  However, the HCL project itself now includes an
[`hcldec`](https://github.com/hashicorp/hcl/tree/main/cmd/hcldec) command that can convert hcl to json if you give it a
specification. One advantage of hcl2json is it doesn't require a specification, so for use cases where the heuristics are correct
it may be easier to use. However, if you have cases hit the limitations of conversion listed above, `hcldec` is probably a
better fit.

## Installation

### Mac OS

Install with [Homebrew](https://formulae.brew.sh/formula/hcl2json):
```shell
brew install hcl2json
```

### Others

Prebuilt binaries are available on the [releases page](https://github.com/tmccombs/hcl2json/releases). There is also a docker image on [dockerhub](https://hub.docker.com/r/tmccombs/hcl2json).

Alternatively you can build from source (see next section).

## Building

You can build and install `hcl2json` using `go get`. Since `hcl2json` uses Go modules, you will need to run this as
`GO11MODULE=on go get github.com/tmccombs/hcl2json`.

Alternatively, you can clone and build the repository:

```
$ git clone https://github.com/tmccombs/hcl2json
$ cd hcl2json
$ go build
```

This will build an `hcl2json` executable in the directory.
