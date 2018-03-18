package outline_test

import (
	"encoding/json"
	"testing"

	"github.com/pokstad/gomate/outline"
)

const expectedOutline = `[{"label":"outline","type":"package","start":78,"end":2584,"children":[{"label":"\"bytes\"","type":"import","start":105,"end":112},{"label":"\"fmt\"","type":"import","start":114,"end":119},{"label":"\"go/ast\"","type":"import","start":121,"end":129},{"label":"\"go/format\"","type":"import","start":131,"end":142},{"label":"\"go/parser\"","type":"import","start":144,"end":155},{"label":"\"go/token\"","type":"import","start":157,"end":167},{"label":"Decl","type":"type","start":243,"end":519},{"label":"ParseFile","type":"function","start":584,"end":2314},{"label":"getReceiverType","type":"function","start":2316,"end":2584}]}]`

func TestParseFile(t *testing.T) {
	d, err := outline.ParseFile("outline.go")
	if err != nil {
		t.Fatalf("unable to parse declarations: %s", err)
	}

	out, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("unable to marshal declarations into JSON: %s", err)
	}

	t.Logf("outline: %s", string(out))

	if string(out) != expectedOutline {
		t.Logf("expected: %s", expectedOutline)
		t.Fail()
	}
}
