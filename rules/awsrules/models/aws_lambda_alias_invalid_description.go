// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsLambdaAliasInvalidDescriptionRule checks the pattern is valid
type AwsLambdaAliasInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsLambdaAliasInvalidDescriptionRule returns new rule with default attributes
func NewAwsLambdaAliasInvalidDescriptionRule() *AwsLambdaAliasInvalidDescriptionRule {
	return &AwsLambdaAliasInvalidDescriptionRule{
		resourceType:  "aws_lambda_alias",
		attributeName: "description",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsLambdaAliasInvalidDescriptionRule) Name() string {
	return "aws_lambda_alias_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaAliasInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaAliasInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaAliasInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaAliasInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
