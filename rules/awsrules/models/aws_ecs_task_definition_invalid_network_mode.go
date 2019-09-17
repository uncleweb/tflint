// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsEcsTaskDefinitionInvalidNetworkModeRule checks the pattern is valid
type AwsEcsTaskDefinitionInvalidNetworkModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEcsTaskDefinitionInvalidNetworkModeRule returns new rule with default attributes
func NewAwsEcsTaskDefinitionInvalidNetworkModeRule() *AwsEcsTaskDefinitionInvalidNetworkModeRule {
	return &AwsEcsTaskDefinitionInvalidNetworkModeRule{
		resourceType:  "aws_ecs_task_definition",
		attributeName: "network_mode",
		enum: []string{
			"bridge",
			"host",
			"awsvpc",
			"none",
		},
	}
}

// Name returns the rule name
func (r *AwsEcsTaskDefinitionInvalidNetworkModeRule) Name() string {
	return "aws_ecs_task_definition_invalid_network_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcsTaskDefinitionInvalidNetworkModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcsTaskDefinitionInvalidNetworkModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcsTaskDefinitionInvalidNetworkModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcsTaskDefinitionInvalidNetworkModeRule) Check(runner *tflint.Runner) error {
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
					`network_mode is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
