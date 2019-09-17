// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/uncleweb/tflint/tflint"
)

func Test_AwsDirectoryServiceDirectoryInvalidTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_directory_service_directory" "foo" {
	type = "ActiveDirectory"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsDirectoryServiceDirectoryInvalidTypeRule(),
					Message: `type is not a valid value`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_directory_service_directory" "foo" {
	type = "SimpleAD"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsDirectoryServiceDirectoryInvalidTypeRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
