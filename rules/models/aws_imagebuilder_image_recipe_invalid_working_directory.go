// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderImageRecipeInvalidWorkingDirectoryRule checks the pattern is valid
type AwsImagebuilderImageRecipeInvalidWorkingDirectoryRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsImagebuilderImageRecipeInvalidWorkingDirectoryRule returns new rule with default attributes
func NewAwsImagebuilderImageRecipeInvalidWorkingDirectoryRule() *AwsImagebuilderImageRecipeInvalidWorkingDirectoryRule {
	return &AwsImagebuilderImageRecipeInvalidWorkingDirectoryRule{
		resourceType:  "aws_imagebuilder_image_recipe",
		attributeName: "working_directory",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsImagebuilderImageRecipeInvalidWorkingDirectoryRule) Name() string {
	return "aws_imagebuilder_image_recipe_invalid_working_directory"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderImageRecipeInvalidWorkingDirectoryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderImageRecipeInvalidWorkingDirectoryRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderImageRecipeInvalidWorkingDirectoryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderImageRecipeInvalidWorkingDirectoryRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"working_directory must be 1024 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"working_directory must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
