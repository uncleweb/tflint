// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsSsmMaintenanceWindowInvalidScheduleRule checks the pattern is valid
type AwsSsmMaintenanceWindowInvalidScheduleRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSsmMaintenanceWindowInvalidScheduleRule returns new rule with default attributes
func NewAwsSsmMaintenanceWindowInvalidScheduleRule() *AwsSsmMaintenanceWindowInvalidScheduleRule {
	return &AwsSsmMaintenanceWindowInvalidScheduleRule{
		resourceType:  "aws_ssm_maintenance_window",
		attributeName: "schedule",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSsmMaintenanceWindowInvalidScheduleRule) Name() string {
	return "aws_ssm_maintenance_window_invalid_schedule"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmMaintenanceWindowInvalidScheduleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmMaintenanceWindowInvalidScheduleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmMaintenanceWindowInvalidScheduleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmMaintenanceWindowInvalidScheduleRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"schedule must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"schedule must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
