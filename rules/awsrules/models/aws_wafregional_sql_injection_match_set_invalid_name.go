// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsWafregionalSQLInjectionMatchSetInvalidNameRule checks the pattern is valid
type AwsWafregionalSQLInjectionMatchSetInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsWafregionalSQLInjectionMatchSetInvalidNameRule returns new rule with default attributes
func NewAwsWafregionalSQLInjectionMatchSetInvalidNameRule() *AwsWafregionalSQLInjectionMatchSetInvalidNameRule {
	return &AwsWafregionalSQLInjectionMatchSetInvalidNameRule{
		resourceType:  "aws_wafregional_sql_injection_match_set",
		attributeName: "name",
		max:           128,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsWafregionalSQLInjectionMatchSetInvalidNameRule) Name() string {
	return "aws_wafregional_sql_injection_match_set_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafregionalSQLInjectionMatchSetInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafregionalSQLInjectionMatchSetInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafregionalSQLInjectionMatchSetInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafregionalSQLInjectionMatchSetInvalidNameRule) Check(runner *tflint.Runner) error {
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
