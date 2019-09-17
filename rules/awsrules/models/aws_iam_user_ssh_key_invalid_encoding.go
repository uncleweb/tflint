// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsIAMUserSSHKeyInvalidEncodingRule checks the pattern is valid
type AwsIAMUserSSHKeyInvalidEncodingRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsIAMUserSSHKeyInvalidEncodingRule returns new rule with default attributes
func NewAwsIAMUserSSHKeyInvalidEncodingRule() *AwsIAMUserSSHKeyInvalidEncodingRule {
	return &AwsIAMUserSSHKeyInvalidEncodingRule{
		resourceType:  "aws_iam_user_ssh_key",
		attributeName: "encoding",
		enum: []string{
			"SSH",
			"PEM",
		},
	}
}

// Name returns the rule name
func (r *AwsIAMUserSSHKeyInvalidEncodingRule) Name() string {
	return "aws_iam_user_ssh_key_invalid_encoding"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMUserSSHKeyInvalidEncodingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMUserSSHKeyInvalidEncodingRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMUserSSHKeyInvalidEncodingRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMUserSSHKeyInvalidEncodingRule) Check(runner *tflint.Runner) error {
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
					`encoding is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
