# 0.6.6

- Use goreleaser to do deploys.

### Warning

This release had some significant changes in how the artifacts were packaged, and doesn't have any meaningful changes to the
package itself. I recommend not using this release.

# 0.6.5

- Update hcl dependency. Should now correctly parse terraform provider functions

# 0.6.4

- Add windows arm64 binary to releases

# 0.6.3

- Properly escape `$${` and `%%{`

# 0.6.2

- Update go-cty and hcl/v2 deps
- Add arm64 docker image

# 0.6.1

- Increase dependency versions

# 0.6.0

- Add -version flag
- Update hcl version

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

