// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppmeshVirtualGatewayInvalidMeshNameRule checks the pattern is valid
type AwsAppmeshVirtualGatewayInvalidMeshNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAppmeshVirtualGatewayInvalidMeshNameRule returns new rule with default attributes
func NewAwsAppmeshVirtualGatewayInvalidMeshNameRule() *AwsAppmeshVirtualGatewayInvalidMeshNameRule {
	return &AwsAppmeshVirtualGatewayInvalidMeshNameRule{
		resourceType:  "aws_appmesh_virtual_gateway",
		attributeName: "mesh_name",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAppmeshVirtualGatewayInvalidMeshNameRule) Name() string {
	return "aws_appmesh_virtual_gateway_invalid_mesh_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppmeshVirtualGatewayInvalidMeshNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppmeshVirtualGatewayInvalidMeshNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppmeshVirtualGatewayInvalidMeshNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppmeshVirtualGatewayInvalidMeshNameRule) Check(runner tflint.Runner) error {
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
					"mesh_name must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"mesh_name must be 1 characters or higher",
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
