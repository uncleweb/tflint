// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsSsmParameterInvalidDescriptionRule checks the pattern is valid
type AwsSsmParameterInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsSsmParameterInvalidDescriptionRule returns new rule with default attributes
func NewAwsSsmParameterInvalidDescriptionRule() *AwsSsmParameterInvalidDescriptionRule {
	return &AwsSsmParameterInvalidDescriptionRule{
		resourceType:  "aws_ssm_parameter",
		attributeName: "description",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsSsmParameterInvalidDescriptionRule) Name() string {
	return "aws_ssm_parameter_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmParameterInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmParameterInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmParameterInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmParameterInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
