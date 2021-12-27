// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderComponentInvalidDescriptionRule checks the pattern is valid
type AwsImagebuilderComponentInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsImagebuilderComponentInvalidDescriptionRule returns new rule with default attributes
func NewAwsImagebuilderComponentInvalidDescriptionRule() *AwsImagebuilderComponentInvalidDescriptionRule {
	return &AwsImagebuilderComponentInvalidDescriptionRule{
		resourceType:  "aws_imagebuilder_component",
		attributeName: "description",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsImagebuilderComponentInvalidDescriptionRule) Name() string {
	return "aws_imagebuilder_component_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderComponentInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderComponentInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderComponentInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderComponentInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
