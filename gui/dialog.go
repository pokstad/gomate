/*Package gui provides Textmate user interface components for requesting user
input. */
package gui

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/mattn/go-plist"
	"github.com/pkg/errors"
	"github.com/pokstad/gomate"
)

type InputDialog struct {
	Title   string // title on top of dialog
	Prompt  string // prompt to show user
	Default string // default value in text input field
}

func (id InputDialog) plistParams() string {
	return fmt.Sprintf(
		`{"title"="%s"; "prompt"="%s"; "string"="%s";}`,
		id.Title, id.Prompt, id.Default,
	)
}

func (id InputDialog) Show(ctx context.Context, env gomate.Env) (string, error) {
	cmd := exec.CommandContext(
		ctx,
		env.Dialog,
		"-cmp",           // center, modal, plist params
		id.plistParams(), // plist params
		filepath.Join(env.SupportPath, "nibs", "RequestString"), // nib file
	)

	output, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "unable to run textmate dialog")
	}

	log.Printf("output from dialog: %s", string(output))

	v, err := plist.Read(bytes.NewBuffer(output))
	if err != nil {
		return "", errors.Wrap(err, "unable to read plist contents")
	}

	str, ok := v.(plist.Dict)["string"].(string)
	if !ok {
		return "", errors.Errorf("unable to obtain response: %#v", str)
	}

	return str, nil

}
