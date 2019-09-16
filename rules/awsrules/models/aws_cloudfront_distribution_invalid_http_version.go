// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsCloudfrontDistributionInvalidHTTPVersionRule checks the pattern is valid
type AwsCloudfrontDistributionInvalidHTTPVersionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCloudfrontDistributionInvalidHTTPVersionRule returns new rule with default attributes
func NewAwsCloudfrontDistributionInvalidHTTPVersionRule() *AwsCloudfrontDistributionInvalidHTTPVersionRule {
	return &AwsCloudfrontDistributionInvalidHTTPVersionRule{
		resourceType:  "aws_cloudfront_distribution",
		attributeName: "http_version",
		enum: []string{
			"http1.1",
			"http2",
		},
	}
}

// Name returns the rule name
func (r *AwsCloudfrontDistributionInvalidHTTPVersionRule) Name() string {
	return "aws_cloudfront_distribution_invalid_http_version"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudfrontDistributionInvalidHTTPVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudfrontDistributionInvalidHTTPVersionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudfrontDistributionInvalidHTTPVersionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudfrontDistributionInvalidHTTPVersionRule) Check(runner *tflint.Runner) error {
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
					`http_version is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
