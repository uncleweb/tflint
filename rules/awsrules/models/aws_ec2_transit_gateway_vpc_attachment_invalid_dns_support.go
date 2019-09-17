// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

// AwsEc2TransitGatewayVpcAttachmentInvalidDNSSupportRule checks the pattern is valid
type AwsEc2TransitGatewayVpcAttachmentInvalidDNSSupportRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2TransitGatewayVpcAttachmentInvalidDNSSupportRule returns new rule with default attributes
func NewAwsEc2TransitGatewayVpcAttachmentInvalidDNSSupportRule() *AwsEc2TransitGatewayVpcAttachmentInvalidDNSSupportRule {
	return &AwsEc2TransitGatewayVpcAttachmentInvalidDNSSupportRule{
		resourceType:  "aws_ec2_transit_gateway_vpc_attachment",
		attributeName: "dns_support",
		enum: []string{
			"enable",
			"disable",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2TransitGatewayVpcAttachmentInvalidDNSSupportRule) Name() string {
	return "aws_ec2_transit_gateway_vpc_attachment_invalid_dns_support"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2TransitGatewayVpcAttachmentInvalidDNSSupportRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEc2TransitGatewayVpcAttachmentInvalidDNSSupportRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2TransitGatewayVpcAttachmentInvalidDNSSupportRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2TransitGatewayVpcAttachmentInvalidDNSSupportRule) Check(runner *tflint.Runner) error {
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
					`dns_support is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
