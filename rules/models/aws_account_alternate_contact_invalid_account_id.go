// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAccountAlternateContactInvalidAccountIDRule checks the pattern is valid
type AwsAccountAlternateContactInvalidAccountIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsAccountAlternateContactInvalidAccountIDRule returns new rule with default attributes
func NewAwsAccountAlternateContactInvalidAccountIDRule() *AwsAccountAlternateContactInvalidAccountIDRule {
	return &AwsAccountAlternateContactInvalidAccountIDRule{
		resourceType:  "aws_account_alternate_contact",
		attributeName: "account_id",
		pattern:       regexp.MustCompile(`^\d{12}$`),
	}
}

// Name returns the rule name
func (r *AwsAccountAlternateContactInvalidAccountIDRule) Name() string {
	return "aws_account_alternate_contact_invalid_account_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAccountAlternateContactInvalidAccountIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAccountAlternateContactInvalidAccountIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAccountAlternateContactInvalidAccountIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAccountAlternateContactInvalidAccountIDRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\d{12}$`),
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
