package code_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/code"
)

func TestGoCode(t *testing.T) {
	srcCode := `package code

import "fmt"

func hello() {
	fmt.Printf("hello")
}
`

	choices, err := code.GoCode{}.Predict(
		context.Background(),
		bytes.NewReader([]byte(srcCode)),
		gomate.Env{
			Cursor: gomate.Cursor{
				Doc:   "code.go",
				Line:  3,
				Index: 1,
			},
		},
	)
	if err != nil {
		t.Fatalf("unable to predict choices: %s", err)
	}

	t.Logf("choices returned: %#v", choices)
}

func TestParseGocode(t *testing.T) {
	jsonBytes := []byte(`[
		0,
		[
			{
				"class":"package",
				"package":"",
				"name":"context",
				"type":""
			},
			{
				"class":"package",
				"package":"",
				"name":"gomate",
				"type":""
			},
			{
				"class":"type",
				"package":"",
				"name":"Choice",
				"type":"struct"
			},{
				"class":"type",
				"package":"",
				"name":"Completer",
				"type":"struct"
			},{
				"class":"type",
				"package":"",
				"name":"GoCode",
				"type":"struct"
			},{
				"class":"type",
				"package":"",
				"name":"Predictor",
				"type":"interface"
			}
		]
	]`)
	code.ParseGocode(jsonBytes)
}
