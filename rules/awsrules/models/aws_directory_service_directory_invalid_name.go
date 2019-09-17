// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsDirectoryServiceDirectoryInvalidNameRule checks the pattern is valid
type AwsDirectoryServiceDirectoryInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsDirectoryServiceDirectoryInvalidNameRule returns new rule with default attributes
func NewAwsDirectoryServiceDirectoryInvalidNameRule() *AwsDirectoryServiceDirectoryInvalidNameRule {
	return &AwsDirectoryServiceDirectoryInvalidNameRule{
		resourceType:  "aws_directory_service_directory",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^([a-zA-Z0-9]+[\\.-])+([a-zA-Z0-9])+$`),
	}
}

// Name returns the rule name
func (r *AwsDirectoryServiceDirectoryInvalidNameRule) Name() string {
	return "aws_directory_service_directory_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDirectoryServiceDirectoryInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDirectoryServiceDirectoryInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDirectoryServiceDirectoryInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDirectoryServiceDirectoryInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^([a-zA-Z0-9]+[\\.-])+([a-zA-Z0-9])+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
