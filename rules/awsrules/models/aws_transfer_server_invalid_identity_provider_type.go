// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsTransferServerInvalidIdentityProviderTypeRule checks the pattern is valid
type AwsTransferServerInvalidIdentityProviderTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsTransferServerInvalidIdentityProviderTypeRule returns new rule with default attributes
func NewAwsTransferServerInvalidIdentityProviderTypeRule() *AwsTransferServerInvalidIdentityProviderTypeRule {
	return &AwsTransferServerInvalidIdentityProviderTypeRule{
		resourceType:  "aws_transfer_server",
		attributeName: "identity_provider_type",
		enum: []string{
			"SERVICE_MANAGED",
			"API_GATEWAY",
		},
	}
}

// Name returns the rule name
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Name() string {
	return "aws_transfer_server_invalid_identity_provider_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Check(runner *tflint.Runner) error {
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
					`identity_provider_type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
