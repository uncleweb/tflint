// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule checks the pattern is valid
type AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule returns new rule with default attributes
func NewAwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule() *AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule {
	return &AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule{
		resourceType:  "aws_config_organization_managed_rule",
		attributeName: "tag_key_scope",
		max:           128,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule) Name() string {
	return "aws_config_organization_managed_rule_invalid_tag_key_scope"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"tag_key_scope must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"tag_key_scope must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
