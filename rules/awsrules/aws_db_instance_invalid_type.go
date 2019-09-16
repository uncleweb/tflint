package awsrules

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsDBInstanceInvalidTypeRule checks whether "aws_db_instance" has invalid intance type.
type AwsDBInstanceInvalidTypeRule struct {
	resourceType  string
	attributeName string
	instanceTypes map[string]bool
}

// NewAwsDBInstanceInvalidTypeRule returns new rule with default attributes
func NewAwsDBInstanceInvalidTypeRule() *AwsDBInstanceInvalidTypeRule {
	return &AwsDBInstanceInvalidTypeRule{
		resourceType:  "aws_db_instance",
		attributeName: "instance_class",
		instanceTypes: map[string]bool{
			"db.t3.micro":     true,
			"db.t3.small":     true,
			"db.t3.medium":    true,
			"db.t3.large":     true,
			"db.t3.xlarge":    true,
			"db.t3.2xlarge":   true,
			"db.t2.micro":     true,
			"db.t2.small":     true,
			"db.t2.medium":    true,
			"db.t2.large":     true,
			"db.t2.xlarge":    true,
			"db.t2.2xlarge":   true,
			"db.m5.large":     true,
			"db.m5.xlarge":    true,
			"db.m5.2xlarge":   true,
			"db.m5.4xlarge":   true,
			"db.m5.12xlarge":  true,
			"db.m5.24xlarge":  true,
			"db.m4.large":     true,
			"db.m4.xlarge":    true,
			"db.m4.2xlarge":   true,
			"db.m4.4xlarge":   true,
			"db.m4.10xlarge":  true,
			"db.m4.16xlarge":  true,
			"db.m3.medium":    true,
			"db.m3.large":     true,
			"db.m3.xlarge":    true,
			"db.m3.2xlarge":   true,
			"db.r5.large":     true,
			"db.r5.xlarge":    true,
			"db.r5.2xlarge":   true,
			"db.r5.4xlarge":   true,
			"db.r5.12xlarge":  true,
			"db.r5.24xlarge":  true,
			"db.r4.large":     true,
			"db.r4.xlarge":    true,
			"db.r4.2xlarge":   true,
			"db.r4.4xlarge":   true,
			"db.r4.8xlarge":   true,
			"db.r4.16xlarge":  true,
			"db.r3.large":     true,
			"db.r3.xlarge":    true,
			"db.r3.2xlarge":   true,
			"db.r3.4xlarge":   true,
			"db.r3.8xlarge":   true,
			"db.t1.micro":     true,
			"db.m1.small":     true,
			"db.m1.medium":    true,
			"db.m1.large":     true,
			"db.m1.xlarge":    true,
			"db.m2.xlarge":    true,
			"db.m2.2xlarge":   true,
			"db.m2.4xlarge":   true,
			"db.cr1.8xlarge":  true,
			"db.x1.16xlarge":  true,
			"db.x1.32xlarge":  true,
			"db.x1e.xlarge":   true,
			"db.x1e.2xlarge":  true,
			"db.x1e.4xlarge":  true,
			"db.x1e.8xlarge":  true,
			"db.x1e.16xlarge": true,
			"db.x1e.32xlarge": true,
		},
	}
}

// Name returns the rule name
func (r *AwsDBInstanceInvalidTypeRule) Name() string {
	return "aws_db_instance_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDBInstanceInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDBInstanceInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDBInstanceInvalidTypeRule) Link() string {
	return "https://github.com/wata727/tflint/blob/master/docs/aws_db_instance_invalid_type.md"
}

// Check checks whether "aws_db_instance" has invalid instance type.
func (r *AwsDBInstanceInvalidTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var instanceType string
		err := runner.EvaluateExpr(attribute.Expr, &instanceType)

		return runner.EnsureNoError(err, func() error {
			if !r.instanceTypes[instanceType] {
				runner.EmitIssue(
					r,
					fmt.Sprintf("\"%s\" is invalid instance type.", instanceType),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
