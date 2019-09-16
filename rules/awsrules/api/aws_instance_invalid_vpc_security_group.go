// This file generated by `tools/api-rule-gen/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsInstanceInvalidVpcSecurityGroupRule checks whether attribute value actually exists
type AwsInstanceInvalidVpcSecurityGroupRule struct {
	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsInstanceInvalidVpcSecurityGroupRule returns new rule with default attributes
func NewAwsInstanceInvalidVpcSecurityGroupRule() *AwsInstanceInvalidVpcSecurityGroupRule {
	return &AwsInstanceInvalidVpcSecurityGroupRule{
		resourceType:  "aws_instance",
		attributeName: "vpc_security_group_ids",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsInstanceInvalidVpcSecurityGroupRule) Name() string {
	return "aws_instance_invalid_vpc_security_group"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsInstanceInvalidVpcSecurityGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsInstanceInvalidVpcSecurityGroupRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsInstanceInvalidVpcSecurityGroupRule) Link() string {
	return ""
}

// Check checks whether the attributes are included in the list retrieved by DescribeSecurityGroups
func (r *AwsInstanceInvalidVpcSecurityGroupRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		if !r.dataPrepared {
			log.Print("[DEBUG] invoking DescribeSecurityGroups")
			var err error
			r.data, err = runner.AwsClient.DescribeSecurityGroups()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking DescribeSecurityGroups",
					Cause:   err,
				}
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		return runner.EachStringSliceExprs(attribute.Expr, func(val string, expr hcl.Expression) {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid security group.`, val),
					expr.Range(),
				)
			}
		})
	})
}
