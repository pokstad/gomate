# gomate
Gomate is a set of TextMate CLI tools for working with Go code

# Install

`go get -u github.com/pokstad/gomate`

# Usage

In the Textmate Go bundle, create a new bundle command with the following script:

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
