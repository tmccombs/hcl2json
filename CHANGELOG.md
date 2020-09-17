# 0.3.0
- Create a seperate library package that is usable by other go projects at github.com/tmccombs/hcl2json/convert
- Add a `-simplify` option that will attempt to simplify expressions that don't have variables or unknown functions


# 0.2.2

- Fix a bug where the last parenthesis of an expression was truncated
- Respect unary operators in literal expressions (specifically, negative numbers are now number literals in the resulting JSON)

