// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsElasticsearchDomainPolicyInvalidDomainNameRule checks the pattern is valid
type AwsElasticsearchDomainPolicyInvalidDomainNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsElasticsearchDomainPolicyInvalidDomainNameRule returns new rule with default attributes
func NewAwsElasticsearchDomainPolicyInvalidDomainNameRule() *AwsElasticsearchDomainPolicyInvalidDomainNameRule {
	return &AwsElasticsearchDomainPolicyInvalidDomainNameRule{
		resourceType:  "aws_elasticsearch_domain_policy",
		attributeName: "domain_name",
		max:           28,
		min:           3,
		pattern:       regexp.MustCompile(`^[a-z][a-z0-9\-]+$`),
	}
}

// Name returns the rule name
func (r *AwsElasticsearchDomainPolicyInvalidDomainNameRule) Name() string {
	return "aws_elasticsearch_domain_policy_invalid_domain_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticsearchDomainPolicyInvalidDomainNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticsearchDomainPolicyInvalidDomainNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticsearchDomainPolicyInvalidDomainNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticsearchDomainPolicyInvalidDomainNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"domain_name must be 28 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"domain_name must be 3 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`domain_name does not match valid pattern ^[a-z][a-z0-9\-]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
