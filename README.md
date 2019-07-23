# hcl2json

This is a tool to convert from [HCL](https://github.com/hashicorp/hcl2/blob/master/hcl/hclsyntax/spec.md) to json, to make it easier for non-go applications and scripts to process HCL inputs (such as terraform config).

If passed the `-pack` option, it converts to the JSON pack format used by [hclpack]( https://godoc.org/github.com/hashicorp/hcl2/hclpack), which contains the original structure of the document. However, all it has for expressions is the source text, which isn't always very useful.

If no options are passed, it converts the provide native HCL file to an (almost) equivalent HCL JSON file. Note, however, that there are some corner cases where it may not be exactly equivalent, especially if the target application makes use of [static analysis](https://github.com/hashicorp/hcl2/blob/master/hcl/hclsyntax/spec.md#static-analysis).
