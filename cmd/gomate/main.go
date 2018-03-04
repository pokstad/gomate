package main

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/guru"
	"github.com/pokstad/gomate/html"
	"github.com/pokstad/gomate/outline"

	"golang.org/x/sync/errgroup"
)

var remarkdownCSS = MustAsset("assets/remarkdown.css")

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("must provide subcommand")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env, err := gomate.LoadEnvironment()
	if err != nil {
		log.Fatalf("unable to load environment vars: %s", err)
	}

	switch os.Args[1] {

	case "install":
		// run deps in parallel to speed up installation
		eg, ctx := errgroup.WithContext(ctx)

		for _, cmd := range []*exec.Cmd{
			exec.CommandContext(ctx, "go", "get", "golang.org/x/tools/cmd/guru"),
		} {
			eg.Go(cmd.Run)
		}

		// wait for all commands to finish running
		if err := eg.Wait(); err != nil {
			log.Fatalf("unable to install dependency: %s", err)
		}

	case "outline":
		decs, err := outline.ParseFile(env)
		checkErr(err)

		err = html.OutlineHTML(os.Stdout, env, decs, remarkdownCSS)
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
