# GoMate

GoMate is a set of TextMate CLI tools for working with Go code. Inspired by 
[syscrusher/golang.tmbundle](https://github.com/syscrusher/golang.tmbundle).

# Install

To get the gomate CLI:

`go get -u github.com/pokstad/gomate`

Then, install the tool's dependencies:

`$GOPATH/bin/gomate install`

# Bundle Install

Until the bundle install is automated, the following needs to be done manually

In the Textmate Go bundle, create a new bundle command with the following 
script:

```
#!/bin/bash
[[ -f "${TM_SUPPORT_PATH}/lib/bash_init.sh" ]] && . "${TM_SUPPORT_PATH}/lib/bash_init.sh"

gomate -path ${TM_FILEPATH} -line ${TM_LINE_NUMBER} -column ${TM_LINE_INDEX}
```

Also, make sure the following options are selected:

- `Input:` is set to `Selection`
- `Format` is set to `Text`
- `Output:` is set to `Show in New Window`
- `Format:` is `HTML`
- `Caret Placement:` is set to `Character Interpolation`

# Roadmap

- Style webpages using [ReMarkdown CSS](https://fvsch.github.io/remarkdown/)
- Reimplement all features of
	[syscrusher/golang.tmbundle](https://github.com/syscrusher/golang.tmbundle)
