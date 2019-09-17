// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsLambdaFunctionInvalidHandlerRule checks the pattern is valid
type AwsLambdaFunctionInvalidHandlerRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsLambdaFunctionInvalidHandlerRule returns new rule with default attributes
func NewAwsLambdaFunctionInvalidHandlerRule() *AwsLambdaFunctionInvalidHandlerRule {
	return &AwsLambdaFunctionInvalidHandlerRule{
		resourceType:  "aws_lambda_function",
		attributeName: "handler",
		max:           128,
		pattern:       regexp.MustCompile(`^[^\s]+$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaFunctionInvalidHandlerRule) Name() string {
	return "aws_lambda_function_invalid_handler"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaFunctionInvalidHandlerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaFunctionInvalidHandlerRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaFunctionInvalidHandlerRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaFunctionInvalidHandlerRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"handler must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`handler does not match valid pattern ^[^\s]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
