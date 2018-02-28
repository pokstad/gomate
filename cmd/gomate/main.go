package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/guru"
	"github.com/pokstad/gomate/outline"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env, err := gomate.LoadEnvironment()
	if err != nil {
		log.Fatalf("unable to load environment vars: %s", err)
	}

	switch os.Args[1] {

	case "outline":
		decs, err := outline.ParseFile(env)
		checkErr(err)

		err = json.NewEncoder(os.Stdout).Encode(decs)
		checkErr(err)

	case "references":
		refs, err := guru.ParseReferrers(ctx, env)
		checkErr(err)

		err = gomate.RenderHTML(os.Stdout, nil, nil, refs)
		checkErr(err)
	}

}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("fatal: %+v", err)
	}
}
