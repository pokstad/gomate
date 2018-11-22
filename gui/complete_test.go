package gui_test

import (
	"context"
	"testing"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/gui"
)

// TestCompleteDialog verifies that the correct commands are run to show a
// completion dialog with images
func TestCompleteDialog(t *testing.T) {
	env := gomate.Env{
		Dialog: "testdata/mock_dialog_1.sh",
	}

	choice, err := gui.CompleteDialog{
		Choices: []gui.Choice{
			{
				Display: "choice 1;",
				Image:   "pkg",
				Insert:  "you picked choice 1",
				Match:   "choice",
			},
		},
		Images: []gui.ChoiceImage{
			{
				Name: "test",
				Path: "/tmp/test.png",
			},
			{
				Name: "test2",
				Path: "/tmp/test2.pdf",
			},
		},
	}.Show(context.Background(), env)
	if err != nil {
		t.Fatalf("Unable to show dialog: %s", err)
	}

	// expected output is a concatenation of the match and insert values
	expectedChoice := "choice 1;you picked choice 1"

	if choice != expectedChoice {
		t.Fatalf("popup dialog returned unexpected output: %s", choice)
	}

	t.Logf("Choice made: %s", choice)
}
