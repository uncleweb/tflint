// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/uncleweb/tflint/tflint"
)

func Test_AwsPlacementGroupInvalidStrategyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_placement_group" "foo" {
	strategy = "instance"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsPlacementGroupInvalidStrategyRule(),
					Message: `strategy is not a valid value`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_placement_group" "foo" {
	strategy = "cluster"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsPlacementGroupInvalidStrategyRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
