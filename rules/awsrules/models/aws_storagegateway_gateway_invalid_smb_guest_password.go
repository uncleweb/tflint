// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule checks the pattern is valid
type AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsStoragegatewayGatewayInvalidSmbGuestPasswordRule returns new rule with default attributes
func NewAwsStoragegatewayGatewayInvalidSmbGuestPasswordRule() *AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule {
	return &AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule{
		resourceType:  "aws_storagegateway_gateway",
		attributeName: "smb_guest_password",
		max:           512,
		min:           6,
		pattern:       regexp.MustCompile(`^[ -~]+$`),
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule) Name() string {
	return "aws_storagegateway_gateway_invalid_smb_guest_password"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayGatewayInvalidSmbGuestPasswordRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"smb_guest_password must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"smb_guest_password must be 6 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`smb_guest_password does not match valid pattern ^[ -~]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
