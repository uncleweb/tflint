// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsIAMGroupPolicyAttachmentInvalidGroupRule checks the pattern is valid
type AwsIAMGroupPolicyAttachmentInvalidGroupRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMGroupPolicyAttachmentInvalidGroupRule returns new rule with default attributes
func NewAwsIAMGroupPolicyAttachmentInvalidGroupRule() *AwsIAMGroupPolicyAttachmentInvalidGroupRule {
	return &AwsIAMGroupPolicyAttachmentInvalidGroupRule{
		resourceType:  "aws_iam_group_policy_attachment",
		attributeName: "group",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMGroupPolicyAttachmentInvalidGroupRule) Name() string {
	return "aws_iam_group_policy_attachment_invalid_group"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMGroupPolicyAttachmentInvalidGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMGroupPolicyAttachmentInvalidGroupRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMGroupPolicyAttachmentInvalidGroupRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMGroupPolicyAttachmentInvalidGroupRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"group must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"group must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`group does not match valid pattern ^[\w+=,.@-]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
