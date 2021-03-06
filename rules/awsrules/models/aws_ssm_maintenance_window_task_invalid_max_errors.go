// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSsmMaintenanceWindowTaskInvalidMaxErrorsRule checks the pattern is valid
type AwsSsmMaintenanceWindowTaskInvalidMaxErrorsRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsmMaintenanceWindowTaskInvalidMaxErrorsRule returns new rule with default attributes
func NewAwsSsmMaintenanceWindowTaskInvalidMaxErrorsRule() *AwsSsmMaintenanceWindowTaskInvalidMaxErrorsRule {
	return &AwsSsmMaintenanceWindowTaskInvalidMaxErrorsRule{
		resourceType:  "aws_ssm_maintenance_window_task",
		attributeName: "max_errors",
		max:           7,
		min:           1,
		pattern:       regexp.MustCompile(`^([1-9][0-9]*|[0]|[1-9][0-9]%|[0-9]%|100%)$`),
	}
}

// Name returns the rule name
func (r *AwsSsmMaintenanceWindowTaskInvalidMaxErrorsRule) Name() string {
	return "aws_ssm_maintenance_window_task_invalid_max_errors"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmMaintenanceWindowTaskInvalidMaxErrorsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmMaintenanceWindowTaskInvalidMaxErrorsRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmMaintenanceWindowTaskInvalidMaxErrorsRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmMaintenanceWindowTaskInvalidMaxErrorsRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"max_errors must be 7 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"max_errors must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([1-9][0-9]*|[0]|[1-9][0-9]%|[0-9]%|100%)$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
