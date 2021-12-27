// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderImagePipelineInvalidDistributionConfigurationArnRule checks the pattern is valid
type AwsImagebuilderImagePipelineInvalidDistributionConfigurationArnRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsImagebuilderImagePipelineInvalidDistributionConfigurationArnRule returns new rule with default attributes
func NewAwsImagebuilderImagePipelineInvalidDistributionConfigurationArnRule() *AwsImagebuilderImagePipelineInvalidDistributionConfigurationArnRule {
	return &AwsImagebuilderImagePipelineInvalidDistributionConfigurationArnRule{
		resourceType:  "aws_imagebuilder_image_pipeline",
		attributeName: "distribution_configuration_arn",
		pattern:       regexp.MustCompile(`^arn:aws[^:]*:imagebuilder:[^:]+:(?:[0-9]{12}|aws):distribution-configuration/[a-z0-9-_]+$`),
	}
}

// Name returns the rule name
func (r *AwsImagebuilderImagePipelineInvalidDistributionConfigurationArnRule) Name() string {
	return "aws_imagebuilder_image_pipeline_invalid_distribution_configuration_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderImagePipelineInvalidDistributionConfigurationArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderImagePipelineInvalidDistributionConfigurationArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderImagePipelineInvalidDistributionConfigurationArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderImagePipelineInvalidDistributionConfigurationArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws[^:]*:imagebuilder:[^:]+:(?:[0-9]{12}|aws):distribution-configuration/[a-z0-9-_]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
