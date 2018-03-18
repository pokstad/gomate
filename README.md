[![CircleCI](https://circleci.com/gh/pokstad/gomate.svg?style=svg)](https://circleci.com/gh/pokstad/gomate)

# GoMate

GoMate is a set of TextMate CLI tools for working with Go code. Inspired by 
[syscrusher/golang.tmbundle](https://github.com/syscrusher/golang.tmbundle).

# Install

To get the gomate CLI:

`go get -u github.com/pokstad/gomate`

Then, install the tool's dependencies:

`$GOPATH/bin/gomate install`

# Usage

Until the bundle install is automated, the following needs to be done manually for each bundle command script:

## Guru

To find references to the symbol under the cursor:

```
#!/bin/bash
gomate references
```

- `Output:` is set to `Show in New Window`
- `Format:` is `HTML`

## Outline

To generate an outline of the current source code file:

```
#!/bin/bash
gomate outline
```

- `Input:` is set to `Selection`
- `Format` is set to `Text`
- `Output:` is set to `Show in New Window`
- `Format:` is `HTML`

## Get Documentation

To view HTML documentation of the symbol under the cursor:

```
#!/bin/bash
gomate getdoc
```

- `Input:` is set to `Document`
- `Format` is set to `Text`
- `Output:` is set to `Show in New Window`
- `Format:` is `HTML`

# Roadmap

- Reimplement all features of
	[syscrusher/golang.tmbundle](https://github.com/syscrusher/golang.tmbundle)
