// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsGlobalacceleratorListenerInvalidProtocolRule checks the pattern is valid
type AwsGlobalacceleratorListenerInvalidProtocolRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGlobalacceleratorListenerInvalidProtocolRule returns new rule with default attributes
func NewAwsGlobalacceleratorListenerInvalidProtocolRule() *AwsGlobalacceleratorListenerInvalidProtocolRule {
	return &AwsGlobalacceleratorListenerInvalidProtocolRule{
		resourceType:  "aws_globalaccelerator_listener",
		attributeName: "protocol",
		enum: []string{
			"TCP",
			"UDP",
		},
	}
}

// Name returns the rule name
func (r *AwsGlobalacceleratorListenerInvalidProtocolRule) Name() string {
	return "aws_globalaccelerator_listener_invalid_protocol"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlobalacceleratorListenerInvalidProtocolRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlobalacceleratorListenerInvalidProtocolRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlobalacceleratorListenerInvalidProtocolRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlobalacceleratorListenerInvalidProtocolRule) Check(runner *tflint.Runner) error {
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
					`protocol is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
