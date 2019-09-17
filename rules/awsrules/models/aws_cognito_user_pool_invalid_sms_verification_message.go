// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsCognitoUserPoolInvalidSmsVerificationMessageRule checks the pattern is valid
type AwsCognitoUserPoolInvalidSmsVerificationMessageRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoUserPoolInvalidSmsVerificationMessageRule returns new rule with default attributes
func NewAwsCognitoUserPoolInvalidSmsVerificationMessageRule() *AwsCognitoUserPoolInvalidSmsVerificationMessageRule {
	return &AwsCognitoUserPoolInvalidSmsVerificationMessageRule{
		resourceType:  "aws_cognito_user_pool",
		attributeName: "sms_verification_message",
		max:           140,
		min:           6,
		pattern:       regexp.MustCompile(`^.*\{####\}.*$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoUserPoolInvalidSmsVerificationMessageRule) Name() string {
	return "aws_cognito_user_pool_invalid_sms_verification_message"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserPoolInvalidSmsVerificationMessageRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoUserPoolInvalidSmsVerificationMessageRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserPoolInvalidSmsVerificationMessageRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserPoolInvalidSmsVerificationMessageRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"sms_verification_message must be 140 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"sms_verification_message must be 6 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`sms_verification_message does not match valid pattern ^.*\{####\}.*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
