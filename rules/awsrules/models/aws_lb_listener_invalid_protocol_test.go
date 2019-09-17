// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/uncleweb/tflint/tflint"
)

func Test_AwsLbListenerInvalidProtocolRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_lb_listener" "foo" {
	protocol = "INVALID"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsLbListenerInvalidProtocolRule(),
					Message: `protocol is not a valid value`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_lb_listener" "foo" {
	protocol = "HTTPS"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsLbListenerInvalidProtocolRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
