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
	"github.com/pokstad/gomate/code"
	"github.com/pokstad/gomate/doc"
	"github.com/pokstad/gomate/gui"
	"github.com/pokstad/gomate/guru"
	"github.com/pokstad/gomate/html"
	"github.com/pokstad/gomate/notes"
	"github.com/pokstad/gomate/outline"
	"github.com/pokstad/gomate/rename"

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
		fatal(errors.New("missing subcommand"))
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

		for _, repo := range []string{
			"golang.org/x/tools/cmd/guru",
			"github.com/zmb3/gogetdoc",
			"github.com/nsf/gocode",
			"github.com/rogpeppe/godef",
			"golang.org/x/tools/cmd/godoc",
			"github.com/alecthomas/gometalinter",
			"golang.org/x/lint/golint",
			"golang.org/x/tools/cmd/gorename",
			"golang.org/x/tools/cmd/goimports",
		} {
			eg.Go(exec.CommandContext(ctx, "go", "get", "-u", repo).Run)
		}

		// wait for all commands to finish running
		if err := eg.Wait(); err != nil {
			checkErr(err)
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

	case "notes":
		pkgNotes, err := notes.AllNotes(env.Drawer.TopDir, env.Drawer.TopDir)
		checkErr(err)

		err = html.RenderNotes(os.Stdout, env, pkgNotes, remarkdownCSS)
		checkErr(err)

	case "rename":
		renameDialog := gui.InputDialog{
			Default: env.Cursor.Word,
			Prompt:  fmt.Sprintf(`Rename \'%s\' to what?`, env.Cursor.Word),
			Title:   "gorename",
		}
		err := rename.AtOffset(ctx, renameDialog, env)
		checkErr(err)

	case "complete":
		completer := code.Completer{
			Predictor: code.GoCode{},
		}

		err = completer.Complete(ctx, env, os.Stdin, os.Stdout)
		checkErr(err)
		os.Exit(int(gomate.ExitInsertSnippet))

	}

}

func checkErr(err error) {
	if err != nil {
		fatal(err)
	}
}

func fatal(e error) {
	err := html.ErrorHTML(os.Stdout, e, logBuf.String(), remarkdownCSS)
	if err != nil {
		_, err := fmt.Fprintf(os.Stdout, "unable to render error HTML: %s", err)
		if err != nil {
			log.Fatalf("unable to print error HTML to file: %s", err)
		}
		os.Exit(int(gomate.ExitCreateNewDoc))
	}
	os.Exit(int(gomate.ExitShowHTML))
}
