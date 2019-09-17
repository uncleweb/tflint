// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsGlobalacceleratorListenerInvalidAcceleratorArnRule checks the pattern is valid
type AwsGlobalacceleratorListenerInvalidAcceleratorArnRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsGlobalacceleratorListenerInvalidAcceleratorArnRule returns new rule with default attributes
func NewAwsGlobalacceleratorListenerInvalidAcceleratorArnRule() *AwsGlobalacceleratorListenerInvalidAcceleratorArnRule {
	return &AwsGlobalacceleratorListenerInvalidAcceleratorArnRule{
		resourceType:  "aws_globalaccelerator_listener",
		attributeName: "accelerator_arn",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsGlobalacceleratorListenerInvalidAcceleratorArnRule) Name() string {
	return "aws_globalaccelerator_listener_invalid_accelerator_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlobalacceleratorListenerInvalidAcceleratorArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlobalacceleratorListenerInvalidAcceleratorArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlobalacceleratorListenerInvalidAcceleratorArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlobalacceleratorListenerInvalidAcceleratorArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"accelerator_arn must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
