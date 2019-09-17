// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsRoute53ResolverRuleAssociationInvalidResolverRuleIDRule checks the pattern is valid
type AwsRoute53ResolverRuleAssociationInvalidResolverRuleIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsRoute53ResolverRuleAssociationInvalidResolverRuleIDRule returns new rule with default attributes
func NewAwsRoute53ResolverRuleAssociationInvalidResolverRuleIDRule() *AwsRoute53ResolverRuleAssociationInvalidResolverRuleIDRule {
	return &AwsRoute53ResolverRuleAssociationInvalidResolverRuleIDRule{
		resourceType:  "aws_route53_resolver_rule_association",
		attributeName: "resolver_rule_id",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverRuleAssociationInvalidResolverRuleIDRule) Name() string {
	return "aws_route53_resolver_rule_association_invalid_resolver_rule_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverRuleAssociationInvalidResolverRuleIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ResolverRuleAssociationInvalidResolverRuleIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverRuleAssociationInvalidResolverRuleIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverRuleAssociationInvalidResolverRuleIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"resolver_rule_id must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"resolver_rule_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
