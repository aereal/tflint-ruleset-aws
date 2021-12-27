// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderImageInvalidImageRecipeArnRule checks the pattern is valid
type AwsImagebuilderImageInvalidImageRecipeArnRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsImagebuilderImageInvalidImageRecipeArnRule returns new rule with default attributes
func NewAwsImagebuilderImageInvalidImageRecipeArnRule() *AwsImagebuilderImageInvalidImageRecipeArnRule {
	return &AwsImagebuilderImageInvalidImageRecipeArnRule{
		resourceType:  "aws_imagebuilder_image",
		attributeName: "image_recipe_arn",
		pattern:       regexp.MustCompile(`^arn:aws[^:]*:imagebuilder:[^:]+:(?:[0-9]{12}|aws):image-recipe/[a-z0-9-_]+/[0-9]+\.[0-9]+\.[0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsImagebuilderImageInvalidImageRecipeArnRule) Name() string {
	return "aws_imagebuilder_image_invalid_image_recipe_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderImageInvalidImageRecipeArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderImageInvalidImageRecipeArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderImageInvalidImageRecipeArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderImageInvalidImageRecipeArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws[^:]*:imagebuilder:[^:]+:(?:[0-9]{12}|aws):image-recipe/[a-z0-9-_]+/[0-9]+\.[0-9]+\.[0-9]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
