package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/doc"
	"github.com/pokstad/gomate/guru"
	"github.com/pokstad/gomate/html"
	"github.com/pokstad/gomate/outline"

	"golang.org/x/sync/errgroup"
)

var (
	// ASSET(markdown.css): used for styling the HTML
	remarkdownCSS = MustAsset("assets/remarkdown.css")
	logBuf        = new(bytes.Buffer)
)

func main() {
	log.SetOutput(logBuf)

	if len(os.Args) < 2 {
		Fatal(errors.New("missing subcommand"))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env, err := gomate.LoadEnvironment()
	if err != nil {
		log.Fatalf("unable to load environment vars: %s", err)
	}

	switch os.Args[1] {

	case "install":
		// install deps in parallel to speed up installation
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
		decs, err := outline.ParseFile(env.Cursor.Doc)
		checkErr(err)

		err = html.OutlineHTML(os.Stdout, env, decs, remarkdownCSS)
		checkErr(err)

	case "references":
		g, err := guru.ObtainGuru(env)
		checkErr(err)

		refs, err := g.ParseReferrers(ctx, env.Cursor, env.Drawer)
		checkErr(err)

		err = html.CodeRefsHTML(os.Stdout, "References", refs, remarkdownCSS)
		checkErr(err)

	case "getdoc":
		getter, err := doc.ObtainGetter(env)
		checkErr(err)

		symbol, err := getter.LookupSymbol(ctx, env.Cursor)
		checkErr(err)

		err = html.SymbolDocHTML(os.Stdout, env.Drawer.TopDir, symbol, remarkdownCSS)
		checkErr(err)

	}

}

func checkErr(err error) {
	if err != nil {
		Fatal(err)
	}
}

func Fatal(e error) {
	err := html.ErrorHTML(os.Stdout, e, logBuf.String(), remarkdownCSS)
	if err != nil {
		fmt.Fprintf(os.Stdout, "unable to render error HTML: %s", err)
		os.Exit(int(gomate.ExitCreateNewDoc))
	}
	os.Exit(int(gomate.ExitShowHTML))
}
