// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule checks the pattern is valid
type AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule returns new rule with default attributes
func NewAwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule() *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule {
	return &AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule{
		resourceType:  "aws_ec2_capacity_reservation",
		attributeName: "instance_match_criteria",
		enum: []string{
			"open",
			"targeted",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Name() string {
	return "aws_ec2_capacity_reservation_invalid_instance_match_criteria"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`instance_match_criteria is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
