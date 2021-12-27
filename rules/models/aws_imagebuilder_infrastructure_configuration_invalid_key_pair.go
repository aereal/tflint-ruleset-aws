// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderInfrastructureConfigurationInvalidKeyPairRule checks the pattern is valid
type AwsImagebuilderInfrastructureConfigurationInvalidKeyPairRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsImagebuilderInfrastructureConfigurationInvalidKeyPairRule returns new rule with default attributes
func NewAwsImagebuilderInfrastructureConfigurationInvalidKeyPairRule() *AwsImagebuilderInfrastructureConfigurationInvalidKeyPairRule {
	return &AwsImagebuilderInfrastructureConfigurationInvalidKeyPairRule{
		resourceType:  "aws_imagebuilder_infrastructure_configuration",
		attributeName: "key_pair",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsImagebuilderInfrastructureConfigurationInvalidKeyPairRule) Name() string {
	return "aws_imagebuilder_infrastructure_configuration_invalid_key_pair"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderInfrastructureConfigurationInvalidKeyPairRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderInfrastructureConfigurationInvalidKeyPairRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderInfrastructureConfigurationInvalidKeyPairRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderInfrastructureConfigurationInvalidKeyPairRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"key_pair must be 1024 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"key_pair must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
