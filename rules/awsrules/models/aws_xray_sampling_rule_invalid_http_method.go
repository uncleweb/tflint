// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsXraySamplingRuleInvalidHTTPMethodRule checks the pattern is valid
type AwsXraySamplingRuleInvalidHTTPMethodRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsXraySamplingRuleInvalidHTTPMethodRule returns new rule with default attributes
func NewAwsXraySamplingRuleInvalidHTTPMethodRule() *AwsXraySamplingRuleInvalidHTTPMethodRule {
	return &AwsXraySamplingRuleInvalidHTTPMethodRule{
		resourceType:  "aws_xray_sampling_rule",
		attributeName: "http_method",
		max:           10,
	}
}

// Name returns the rule name
func (r *AwsXraySamplingRuleInvalidHTTPMethodRule) Name() string {
	return "aws_xray_sampling_rule_invalid_http_method"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsXraySamplingRuleInvalidHTTPMethodRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsXraySamplingRuleInvalidHTTPMethodRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsXraySamplingRuleInvalidHTTPMethodRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsXraySamplingRuleInvalidHTTPMethodRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"http_method must be 10 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
