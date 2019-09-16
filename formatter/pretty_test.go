package formatter

import (
	"bytes"
	"errors"
	"testing"

	"github.com/fatih/color"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

func Test_prettyPrint(t *testing.T) {
	// Disable color
	color.NoColor = true

	cases := []struct {
		Name    string
		Issues  tflint.Issues
		Error   *tflint.Error
		Sources map[string][]byte
		Stdout  string
		Stderr  string
	}{
		{
			Name:   "no issues",
			Issues: tflint.Issues{},
			Stdout: "",
		},
		{
			Name: "issues",
			Issues: tflint.Issues{
				{
					Rule:    &testRule{},
					Message: "test",
					Range: hcl.Range{
						Filename: "test.tf",
						Start:    hcl.Pos{Line: 1, Column: 1, Byte: 0},
						End:      hcl.Pos{Line: 1, Column: 4, Byte: 3},
					},
					Callers: []hcl.Range{
						{
							Filename: "test.tf",
							Start:    hcl.Pos{Line: 1, Column: 1, Byte: 0},
							End:      hcl.Pos{Line: 1, Column: 4, Byte: 3},
						},
						{
							Filename: "module.tf",
							Start:    hcl.Pos{Line: 2, Column: 3, Byte: 0},
							End:      hcl.Pos{Line: 2, Column: 6, Byte: 3},
						},
					},
				},
			},
			Sources: map[string][]byte{
				"test.tf": []byte("foo = 1"),
			},
			Stdout: `1 issue(s) found:

Error: test (test_rule)

  on test.tf line 1:
   1: foo = 1

Callers:
   test.tf:1,1-4
   module.tf:2,3-6

Reference: https://github.com

`,
		},
		{
			Name: "no sources",
			Issues: tflint.Issues{
				{
					Rule:    &testRule{},
					Message: "test",
					Range: hcl.Range{
						Filename: "test.tf",
						Start:    hcl.Pos{Line: 1, Column: 1, Byte: 0},
						End:      hcl.Pos{Line: 1, Column: 4, Byte: 3},
					},
				},
			},
			Stdout: `1 issue(s) found:

Error: test (test_rule)

  on test.tf line 1:
   (source code not available)

Reference: https://github.com

`,
		},
		{
			Name:   "error",
			Issues: tflint.Issues{},
			Error:  tflint.NewContextError("Failed to work", errors.New("I don't feel like working")),
			Stderr: `Failed to work. An error occurred:

Error: I don't feel like working

`,
		},
	}

	for _, tc := range cases {
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}
		formatter := &Formatter{Stdout: stdout, Stderr: stderr}

		formatter.prettyPrint(tc.Issues, tc.Error, tc.Sources)

		if stdout.String() != tc.Stdout {
			t.Fatalf("Failed %s test: expected=%s, stdout=%s", tc.Name, tc.Stdout, stdout.String())
		}
		if stderr.String() != tc.Stderr {
			t.Fatalf("Failed %s test: expected=%s, stderr=%s", tc.Name, tc.Stderr, stderr.String())
		}
	}
}
