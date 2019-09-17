// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsKmsGrantInvalidRetiringPrincipalRule checks the pattern is valid
type AwsKmsGrantInvalidRetiringPrincipalRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsKmsGrantInvalidRetiringPrincipalRule returns new rule with default attributes
func NewAwsKmsGrantInvalidRetiringPrincipalRule() *AwsKmsGrantInvalidRetiringPrincipalRule {
	return &AwsKmsGrantInvalidRetiringPrincipalRule{
		resourceType:  "aws_kms_grant",
		attributeName: "retiring_principal",
		max:           256,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@:/-]+$`),
	}
}

// Name returns the rule name
func (r *AwsKmsGrantInvalidRetiringPrincipalRule) Name() string {
	return "aws_kms_grant_invalid_retiring_principal"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKmsGrantInvalidRetiringPrincipalRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKmsGrantInvalidRetiringPrincipalRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKmsGrantInvalidRetiringPrincipalRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKmsGrantInvalidRetiringPrincipalRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"retiring_principal must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"retiring_principal must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`retiring_principal does not match valid pattern ^[\w+=,.@:/-]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
