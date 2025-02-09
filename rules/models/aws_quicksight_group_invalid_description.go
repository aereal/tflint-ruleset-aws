// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsQuicksightGroupInvalidDescriptionRule checks the pattern is valid
type AwsQuicksightGroupInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsQuicksightGroupInvalidDescriptionRule returns new rule with default attributes
func NewAwsQuicksightGroupInvalidDescriptionRule() *AwsQuicksightGroupInvalidDescriptionRule {
	return &AwsQuicksightGroupInvalidDescriptionRule{
		resourceType:  "aws_quicksight_group",
		attributeName: "description",
		max:           512,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsQuicksightGroupInvalidDescriptionRule) Name() string {
	return "aws_quicksight_group_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsQuicksightGroupInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsQuicksightGroupInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsQuicksightGroupInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsQuicksightGroupInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"description must be 1 characters or higher",
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
