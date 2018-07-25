[![CircleCI](https://circleci.com/gh/pokstad/gomate.svg?style=svg)](https://circleci.com/gh/pokstad/gomate)

# Gomate

Gomate is a set of TextMate CLI tools for working with Go code. Inspired by 
[syscrusher/golang.tmbundle](https://github.com/syscrusher/golang.tmbundle).

Gomate embraces the Unix philosophy adopted by Textmate by utilizing simple CLI tool constructs, such as:

- Command line arguments to indicate desired action
- Sourcing environment variable for context of operation
- Reading input from STDIN and producing desired output on STDOUT

Additionally, Gomate is comprised of many smaller packages with dedicated functions to allow for maximum reuse outside the scope of Textmate. For example, you may want to use the note parsing package and build your own extension to a different tool (e.g. VSCode or VIM).

The modular design also encouraged the development of some dedicated CLI tools:

- [notes](notes/notes) - provided a path, will scan recursively and return all godoc notes in all packages
  - Installation: `go get github.com/pokstad/gomate/notes/notes`
  - Usage: `$GOPATH/bin/notes [OPTIONS] PACKAGE_PATH`

## Install

To get the gomate CLI:

`go get -u github.com/pokstad/gomate`

Then, install the tool's dependencies:

`$GOPATH/bin/gomate install`

## Usage

Until the bundle install is automated, the following needs to be done manually for each bundle command script:

### References

To find references to the symbol under the cursor:

```
#!/bin/bash
gomate references
```

- `Output:` is set to `Show in New Window`
- `Format:` is `HTML`

External dependencies: [guru](https://golang.org/x/tools/cmd/guru)

### Outline

To generate an outline of the current source code file:

```
#!/bin/bash
gomate outline
```

- `Input:` is set to `Selection`
- `Format` is set to `Text`
- `Output:` is set to `Show in New Window`
- `Format:` is `HTML`

### Get Documentation

To view HTML documentation of the symbol under the cursor:

```
#!/bin/bash
gomate getdoc
```

- `Input:` is set to `Document`
- `Format` is set to `Text`
- `Output:` is set to `Show in New Window`
- `Format:` is `HTML`

External dependencies: [gogetdoc](https://github.com/zmb3/gogetdoc)

## Roadmap

- Testing coverage and advanced options for individual test functions (similar to VSCode code lense)
- Reimplement all features of existing Textmate plugin
	[syscrusher/golang.tmbundle](https://github.com/syscrusher/golang.tmbundle)
- Web service to support rich contexts that allow interactions via AJAX calls