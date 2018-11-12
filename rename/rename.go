package rename

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/pokstad/gomate"
)

// InputDialog is how we request the new symbol identifier from the user
type InputDialog interface {
	// Show will return a user defined text or error
	Show(context.Context, gomate.Env) (string, error)
}

// AtOffset will rename a symbol at the specifed offset after requesting a new
// symbol name from a dialog shown to the user
func AtOffset(ctx context.Context, id InputDialog, env gomate.Env) error {
	resp, err := id.Show(ctx, env)
	if err != nil {
		return errors.WithStack(err)
	}

	offset, err := env.Cursor.RuneOffset()
	if err != nil {
		return errors.WithStack(err)
	}

	return exec.CommandContext(
		ctx,
		"gorename",
		"-offset", fmt.Sprintf("%s:#%d", env.Cursor.Doc, offset),
		"-to", resp,
	).Run()
}
