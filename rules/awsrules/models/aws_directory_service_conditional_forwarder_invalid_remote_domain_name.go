// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule checks the pattern is valid
type AwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule returns new rule with default attributes
func NewAwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule() *AwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule {
	return &AwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule{
		resourceType:  "aws_directory_service_conditional_forwarder",
		attributeName: "remote_domain_name",
		pattern:       regexp.MustCompile(`^([a-zA-Z0-9]+[\\.-])+([a-zA-Z0-9])+[.]?$`),
	}
}

// Name returns the rule name
func (r *AwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule) Name() string {
	return "aws_directory_service_conditional_forwarder_invalid_remote_domain_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`remote_domain_name does not match valid pattern ^([a-zA-Z0-9]+[\\.-])+([a-zA-Z0-9])+[.]?$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
