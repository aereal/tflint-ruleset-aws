// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderInfrastructureConfigurationInvalidDescriptionRule checks the pattern is valid
type AwsImagebuilderInfrastructureConfigurationInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsImagebuilderInfrastructureConfigurationInvalidDescriptionRule returns new rule with default attributes
func NewAwsImagebuilderInfrastructureConfigurationInvalidDescriptionRule() *AwsImagebuilderInfrastructureConfigurationInvalidDescriptionRule {
	return &AwsImagebuilderInfrastructureConfigurationInvalidDescriptionRule{
		resourceType:  "aws_imagebuilder_infrastructure_configuration",
		attributeName: "description",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsImagebuilderInfrastructureConfigurationInvalidDescriptionRule) Name() string {
	return "aws_imagebuilder_infrastructure_configuration_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderInfrastructureConfigurationInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderInfrastructureConfigurationInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderInfrastructureConfigurationInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderInfrastructureConfigurationInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 1024 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"description must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
