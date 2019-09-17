// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidCertificateRule checks the pattern is valid
type AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidCertificateRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsWorklinkWebsiteCertificateAuthorityAssociationInvalidCertificateRule returns new rule with default attributes
func NewAwsWorklinkWebsiteCertificateAuthorityAssociationInvalidCertificateRule() *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidCertificateRule {
	return &AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidCertificateRule{
		resourceType:  "aws_worklink_website_certificate_authority_association",
		attributeName: "certificate",
		max:           8192,
		min:           1,
		pattern:       regexp.MustCompile(`^-{5}BEGIN CERTIFICATE-{5}\x{000D}?\x{000A}([A-Za-z0-9/+]{64}\x{000D}?\x{000A})*[A-Za-z0-9/+]{1,64}={0,2}\x{000D}?\x{000A}-{5}END CERTIFICATE-{5}(\x{000D}?\x{000A})?$`),
	}
}

// Name returns the rule name
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidCertificateRule) Name() string {
	return "aws_worklink_website_certificate_authority_association_invalid_certificate"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidCertificateRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidCertificateRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidCertificateRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidCertificateRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"certificate must be 8192 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"certificate must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`certificate does not match valid pattern ^-{5}BEGIN CERTIFICATE-{5}\x{000D}?\x{000A}([A-Za-z0-9/+]{64}\x{000D}?\x{000A})*[A-Za-z0-9/+]{1,64}={0,2}\x{000D}?\x{000A}-{5}END CERTIFICATE-{5}(\x{000D}?\x{000A})?$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
