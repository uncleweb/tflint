// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsKmsGrantInvalidKeyIDRule checks the pattern is valid
type AwsKmsGrantInvalidKeyIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsKmsGrantInvalidKeyIDRule returns new rule with default attributes
func NewAwsKmsGrantInvalidKeyIDRule() *AwsKmsGrantInvalidKeyIDRule {
	return &AwsKmsGrantInvalidKeyIDRule{
		resourceType:  "aws_kms_grant",
		attributeName: "key_id",
		max:           2048,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsKmsGrantInvalidKeyIDRule) Name() string {
	return "aws_kms_grant_invalid_key_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKmsGrantInvalidKeyIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKmsGrantInvalidKeyIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKmsGrantInvalidKeyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKmsGrantInvalidKeyIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"key_id must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"key_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
