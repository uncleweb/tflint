// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsAcmCertificateInvalidCertificateBodyRule checks the pattern is valid
type AwsAcmCertificateInvalidCertificateBodyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAcmCertificateInvalidCertificateBodyRule returns new rule with default attributes
func NewAwsAcmCertificateInvalidCertificateBodyRule() *AwsAcmCertificateInvalidCertificateBodyRule {
	return &AwsAcmCertificateInvalidCertificateBodyRule{
		resourceType:  "aws_acm_certificate",
		attributeName: "certificate_body",
		max:           32768,
		min:           1,
		pattern:       regexp.MustCompile(`^-{5}BEGIN CERTIFICATE-{5}\x{000D}?\x{000A}([A-Za-z0-9/+]{64}\x{000D}?\x{000A})*[A-Za-z0-9/+]{1,64}={0,2}\x{000D}?\x{000A}-{5}END CERTIFICATE-{5}(\x{000D}?\x{000A})?$`),
	}
}

// Name returns the rule name
func (r *AwsAcmCertificateInvalidCertificateBodyRule) Name() string {
	return "aws_acm_certificate_invalid_certificate_body"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAcmCertificateInvalidCertificateBodyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAcmCertificateInvalidCertificateBodyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAcmCertificateInvalidCertificateBodyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAcmCertificateInvalidCertificateBodyRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"certificate_body must be 32768 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"certificate_body must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`certificate_body does not match valid pattern ^-{5}BEGIN CERTIFICATE-{5}\x{000D}?\x{000A}([A-Za-z0-9/+]{64}\x{000D}?\x{000A})*[A-Za-z0-9/+]{1,64}={0,2}\x{000D}?\x{000A}-{5}END CERTIFICATE-{5}(\x{000D}?\x{000A})?$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
