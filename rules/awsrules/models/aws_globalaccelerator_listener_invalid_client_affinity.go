// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsGlobalacceleratorListenerInvalidClientAffinityRule checks the pattern is valid
type AwsGlobalacceleratorListenerInvalidClientAffinityRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGlobalacceleratorListenerInvalidClientAffinityRule returns new rule with default attributes
func NewAwsGlobalacceleratorListenerInvalidClientAffinityRule() *AwsGlobalacceleratorListenerInvalidClientAffinityRule {
	return &AwsGlobalacceleratorListenerInvalidClientAffinityRule{
		resourceType:  "aws_globalaccelerator_listener",
		attributeName: "client_affinity",
		enum: []string{
			"NONE",
			"SOURCE_IP",
		},
	}
}

// Name returns the rule name
func (r *AwsGlobalacceleratorListenerInvalidClientAffinityRule) Name() string {
	return "aws_globalaccelerator_listener_invalid_client_affinity"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlobalacceleratorListenerInvalidClientAffinityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlobalacceleratorListenerInvalidClientAffinityRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlobalacceleratorListenerInvalidClientAffinityRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlobalacceleratorListenerInvalidClientAffinityRule) Check(runner *tflint.Runner) error {
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
					`client_affinity is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
