// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/uncleweb/tflint/tflint"
)

func Test_AwsCognitoUserPoolInvalidSmsAuthenticationMessageRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cognito_user_pool" "foo" {
	sms_authentication_message = "Authentication code"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCognitoUserPoolInvalidSmsAuthenticationMessageRule(),
					Message: `sms_authentication_message does not match valid pattern ^.*\{####\}.*$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cognito_user_pool" "foo" {
	sms_authentication_message = "Authentication code is {####}"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCognitoUserPoolInvalidSmsAuthenticationMessageRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
