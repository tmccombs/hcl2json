# 0.5.0

- Make a few functions public in library
- Fix bug where stdin wasn't used if no files are passed as arguments

# 0.4.0
- Update dependencies
- Add support for passing multiple files to hcl2json

# 0.3.6
- Update dependencies
- Add linux/arm64 to releases page

# 0.3.5
- Update dependencies
- Add automated build for docker image
- Increase minimum go version to 1.18

# 0.3.4
- Update dependencies
- Fix panic when objects contain a key of "null"

# 0.3.3
- Export convertFile function
- Update dependencies
- Fix failing to parse if no newline at end of file

# 0.3.0
- Create a seperate library package that is usable by other go projects at github.com/tmccombs/hcl2json/convert
- Add a `-simplify` option that will attempt to simplify expressions that don't have variables or unknown functions


# 0.2.2

- Fix a bug where the last parenthesis of an expression was truncated
- Respect unary operators in literal expressions (specifically, negative numbers are now number literals in the resulting JSON)

