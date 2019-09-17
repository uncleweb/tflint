// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsSsmMaintenanceWindowInvalidNameRule checks the pattern is valid
type AwsSsmMaintenanceWindowInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsmMaintenanceWindowInvalidNameRule returns new rule with default attributes
func NewAwsSsmMaintenanceWindowInvalidNameRule() *AwsSsmMaintenanceWindowInvalidNameRule {
	return &AwsSsmMaintenanceWindowInvalidNameRule{
		resourceType:  "aws_ssm_maintenance_window",
		attributeName: "name",
		max:           128,
		min:           3,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_\-.]{3,128}$`),
	}
}

// Name returns the rule name
func (r *AwsSsmMaintenanceWindowInvalidNameRule) Name() string {
	return "aws_ssm_maintenance_window_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmMaintenanceWindowInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmMaintenanceWindowInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmMaintenanceWindowInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmMaintenanceWindowInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 3 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^[a-zA-Z0-9_\-.]{3,128}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
