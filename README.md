# hcl2json

This is a tool to convert from [HCL](https://github.com/hashicorp/hcl2/blob/master/hcl/hclsyntax/spec.md) to json, to make it easier for non-go applications and scripts to process HCL inputs (such as terraform config).

At the moment, it just converts to the JSON pack format used by [hclpack]( https://godoc.org/github.com/hashicorp/hcl2/hclpack), which contains the structure of the document, all it has for expressions is the source text, which isn't always very useful.
