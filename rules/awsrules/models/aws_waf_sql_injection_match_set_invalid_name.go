// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsWafSQLInjectionMatchSetInvalidNameRule checks the pattern is valid
type AwsWafSQLInjectionMatchSetInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsWafSQLInjectionMatchSetInvalidNameRule returns new rule with default attributes
func NewAwsWafSQLInjectionMatchSetInvalidNameRule() *AwsWafSQLInjectionMatchSetInvalidNameRule {
	return &AwsWafSQLInjectionMatchSetInvalidNameRule{
		resourceType:  "aws_waf_sql_injection_match_set",
		attributeName: "name",
		max:           128,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsWafSQLInjectionMatchSetInvalidNameRule) Name() string {
	return "aws_waf_sql_injection_match_set_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafSQLInjectionMatchSetInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafSQLInjectionMatchSetInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafSQLInjectionMatchSetInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafSQLInjectionMatchSetInvalidNameRule) Check(runner *tflint.Runner) error {
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
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
