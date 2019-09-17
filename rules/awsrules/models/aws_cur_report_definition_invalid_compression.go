// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsCurReportDefinitionInvalidCompressionRule checks the pattern is valid
type AwsCurReportDefinitionInvalidCompressionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCurReportDefinitionInvalidCompressionRule returns new rule with default attributes
func NewAwsCurReportDefinitionInvalidCompressionRule() *AwsCurReportDefinitionInvalidCompressionRule {
	return &AwsCurReportDefinitionInvalidCompressionRule{
		resourceType:  "aws_cur_report_definition",
		attributeName: "compression",
		enum: []string{
			"ZIP",
			"GZIP",
			"Parquet",
		},
	}
}

// Name returns the rule name
func (r *AwsCurReportDefinitionInvalidCompressionRule) Name() string {
	return "aws_cur_report_definition_invalid_compression"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCurReportDefinitionInvalidCompressionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCurReportDefinitionInvalidCompressionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCurReportDefinitionInvalidCompressionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCurReportDefinitionInvalidCompressionRule) Check(runner *tflint.Runner) error {
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
					`compression is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
