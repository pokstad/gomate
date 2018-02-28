package main

import (
	"flag"
	"os"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/outline"
)

var (
	outlineFlags    = flag.NewFlagSet("outline", flag.ExitOnError)
	referencesFlags = flag.NewFlagSet("references", flag.ExitOnError)
)

func main() {
	flag.Parse()
	eVars := gomate.LoadEnvironment()

	switch os.Args[1] {

	case "outline":
		outlineFlags.Parse(os.Args[2:])
		outline.ParseFile(eVars.CurrDoc)

	case "references":
		referencesFlags.Parse(os.Args[2:])

	}

}
