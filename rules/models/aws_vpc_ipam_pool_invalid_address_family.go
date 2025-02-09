// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsVpcIpamPoolInvalidAddressFamilyRule checks the pattern is valid
type AwsVpcIpamPoolInvalidAddressFamilyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsVpcIpamPoolInvalidAddressFamilyRule returns new rule with default attributes
func NewAwsVpcIpamPoolInvalidAddressFamilyRule() *AwsVpcIpamPoolInvalidAddressFamilyRule {
	return &AwsVpcIpamPoolInvalidAddressFamilyRule{
		resourceType:  "aws_vpc_ipam_pool",
		attributeName: "address_family",
		enum: []string{
			"ipv4",
			"ipv6",
		},
	}
}

// Name returns the rule name
func (r *AwsVpcIpamPoolInvalidAddressFamilyRule) Name() string {
	return "aws_vpc_ipam_pool_invalid_address_family"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsVpcIpamPoolInvalidAddressFamilyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsVpcIpamPoolInvalidAddressFamilyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsVpcIpamPoolInvalidAddressFamilyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsVpcIpamPoolInvalidAddressFamilyRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as address_family`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
