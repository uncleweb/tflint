// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/uncleweb/tflint/tflint"
)

func Test_AwsCloudhsmV2HsmInvalidSubnetIDRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudhsm_v2_hsm" "foo" {
	subnet_id = "0e358c43"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudhsmV2HsmInvalidSubnetIDRule(),
					Message: `subnet_id does not match valid pattern ^subnet-[0-9a-fA-F]{8,17}$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudhsm_v2_hsm" "foo" {
	subnet_id = "subnet-0e358c43"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudhsmV2HsmInvalidSubnetIDRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
