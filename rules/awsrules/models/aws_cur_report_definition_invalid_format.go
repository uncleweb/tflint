// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsCurReportDefinitionInvalidFormatRule checks the pattern is valid
type AwsCurReportDefinitionInvalidFormatRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCurReportDefinitionInvalidFormatRule returns new rule with default attributes
func NewAwsCurReportDefinitionInvalidFormatRule() *AwsCurReportDefinitionInvalidFormatRule {
	return &AwsCurReportDefinitionInvalidFormatRule{
		resourceType:  "aws_cur_report_definition",
		attributeName: "format",
		enum: []string{
			"textORcsv",
			"Parquet",
		},
	}
}

// Name returns the rule name
func (r *AwsCurReportDefinitionInvalidFormatRule) Name() string {
	return "aws_cur_report_definition_invalid_format"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCurReportDefinitionInvalidFormatRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCurReportDefinitionInvalidFormatRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCurReportDefinitionInvalidFormatRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCurReportDefinitionInvalidFormatRule) Check(runner *tflint.Runner) error {
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
					`format is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
