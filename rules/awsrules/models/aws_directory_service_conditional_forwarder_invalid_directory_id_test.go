// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/uncleweb/tflint/tflint"
)

func Test_AwsDirectoryServiceConditionalForwarderInvalidDirectoryIDRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_directory_service_conditional_forwarder" "foo" {
	directory_id = "1234567890"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsDirectoryServiceConditionalForwarderInvalidDirectoryIDRule(),
					Message: `directory_id does not match valid pattern ^d-[0-9a-f]{10}$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_directory_service_conditional_forwarder" "foo" {
	directory_id = "d-1234567890"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsDirectoryServiceConditionalForwarderInvalidDirectoryIDRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
