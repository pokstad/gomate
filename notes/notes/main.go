/* notes is a utility that recursively scans all note markings in a directory
and returns a representation of those notes mapped by tag and ID to the actual
note and location. This representation essentially looks like this:

  TAG1:
    ID1: Note (location)
  TAG2:
    ID2: Note (location)

See usage string for details on proper usage.
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pokstad/gomate/notes"
)

const usage = `
  OVERVIEW:

    notes scans Go source code files for notations

  USAGE:

    notes [OPTIONS] TARGET

  TARGET:

    the Go package path you wish to scan

  OPTIONS:

    -output [json|textmate]

`

const (
	outputJSON     = "json"
	outputTextmate = "textmate"
)

var (
	output = flag.String("output", outputJSON, "output format to return")
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}

	target := os.Args[len(os.Args)-1]

	allNotes, err := notes.AllNotes(target, "")
	if err != nil {
		log.Fatalf("unable to scan notes from target %s: %s", target, err)
	}

	print(*output)

	switch *output {

	case outputJSON:
		err := json.NewEncoder(os.Stdout).Encode(allNotes)
		if err != nil {
			log.Fatalf("unable to encode notes to JSON: %s", err)
		}

	case outputTextmate:
		// TODO
		log.Fatal("not implemented yet!")

	default:
		log.Fatalf("unknown output type: %s", *output)

	}
}
