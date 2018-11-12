// +build textmate

// All tests in this file must be executed with Textmate app as the parent
// process or they will fail.

package gui_test

import (
	"context"
	"testing"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/gui"
)

func TestInputDialog(t *testing.T) {
	env, err := gomate.LoadEnvironment()
	if err != nil {
		t.Fatalf("unable to load env vars: %s", err)
	}

	resp, err := gui.InputDialog{
		Title:   "Lazy Sunday",
		Prompt:  "What would you like to do?",
		Default: "nothing",
	}.ShowInputDialog(
		context.Background(),
		env,
	)
	if err != nil {
		t.Fatalf("unable to obtain response from dialog: %s", err)
	}

	t.Logf("resp: %s", resp)
}
