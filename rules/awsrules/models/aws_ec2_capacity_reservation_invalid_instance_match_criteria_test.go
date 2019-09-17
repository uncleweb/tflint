// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/uncleweb/tflint/tflint"
)

func Test_AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ec2_capacity_reservation" "foo" {
	instance_match_criteria = "close"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule(),
					Message: `instance_match_criteria is not a valid value`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ec2_capacity_reservation" "foo" {
	instance_match_criteria = "open"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
