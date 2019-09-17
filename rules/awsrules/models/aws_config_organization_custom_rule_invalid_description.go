// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsConfigOrganizationCustomRuleInvalidDescriptionRule checks the pattern is valid
type AwsConfigOrganizationCustomRuleInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsConfigOrganizationCustomRuleInvalidDescriptionRule returns new rule with default attributes
func NewAwsConfigOrganizationCustomRuleInvalidDescriptionRule() *AwsConfigOrganizationCustomRuleInvalidDescriptionRule {
	return &AwsConfigOrganizationCustomRuleInvalidDescriptionRule{
		resourceType:  "aws_config_organization_custom_rule",
		attributeName: "description",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationCustomRuleInvalidDescriptionRule) Name() string {
	return "aws_config_organization_custom_rule_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationCustomRuleInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationCustomRuleInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationCustomRuleInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationCustomRuleInvalidDescriptionRule) Check(runner *tflint.Runner) error {
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
