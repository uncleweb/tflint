package terraformrules

import (
	"testing"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/uncleweb/tflint/tflint"
)

func Test_TerraformModulePinnedSource(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "git module is not pinned",
			Content: `
module "unpinned" {
  source = "git://hashicorp.com/consul.git"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewTerraformModulePinnedSourceRule(),
					Message: "Module source \"git://hashicorp.com/consul.git\" is not pinned",
					Range: hcl.Range{
						Filename: "module.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 44},
					},
				},
			},
		},
		{
			Name: "git module reference is default",
			Content: `
module "default_git" {
  source = "git://hashicorp.com/consul.git?ref=master"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewTerraformModulePinnedSourceRule(),
					Message: "Module source \"git://hashicorp.com/consul.git?ref=master\" uses default ref \"master\"",
					Range: hcl.Range{
						Filename: "module.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 55},
					},
				},
			},
		},
		{
			Name: "git module reference is pinned",
			Content: `
module "pinned_git" {
  source = "git://hashicorp.com/consul.git?ref=pinned"
}`,
			Expected: tflint.Issues{},
		},
		{
			Name: "github module is not pinned",
			Content: `
module "unpinned" {
  source = "github.com/hashicorp/consul"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewTerraformModulePinnedSourceRule(),
					Message: "Module source \"github.com/hashicorp/consul\" is not pinned",
					Range: hcl.Range{
						Filename: "module.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 41},
					},
				},
			},
		},
		{
			Name: "github module reference is default",
			Content: `
module "default_git" {
  source = "github.com/hashicorp/consul.git?ref=master"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewTerraformModulePinnedSourceRule(),
					Message: "Module source \"github.com/hashicorp/consul.git?ref=master\" uses default ref \"master\"",
					Range: hcl.Range{
						Filename: "module.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 56},
					},
				},
			},
		},
		{
			Name: "github module reference is pinned",
			Content: `
module "pinned_git" {
  source = "github.com/hashicorp/consul.git?ref=pinned"
}`,
			Expected: tflint.Issues{},
		},
		{
			Name: "bitbucket module is not pinned",
			Content: `
module "unpinned" {
  source = "bitbucket.org/hashicorp/consul"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewTerraformModulePinnedSourceRule(),
					Message: "Module source \"bitbucket.org/hashicorp/consul\" is not pinned",
					Range: hcl.Range{
						Filename: "module.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 44},
					},
				},
			},
		},
		{
			Name: "bitbucket module reference is default",
			Content: `
module "default_git" {
  source = "bitbucket.org/hashicorp/consul.git?ref=master"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewTerraformModulePinnedSourceRule(),
					Message: "Module source \"bitbucket.org/hashicorp/consul.git?ref=master\" uses default ref \"master\"",
					Range: hcl.Range{
						Filename: "module.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 59},
					},
				},
			},
		},
		{
			Name: "bitbucket module reference is pinned",
			Content: `
module "pinned_git" {
  source = "bitbucket.org/hashicorp/consul.git?ref=pinned"
}`,
			Expected: tflint.Issues{},
		},
		{
			Name: "generic git (git::https) module reference is not pinned",
			Content: `
module "unpinned_generic_git_https" {
  source = "git::https://hashicorp.com/consul.git"
}
`,
			Expected: tflint.Issues{
				{
					Rule:    NewTerraformModulePinnedSourceRule(),
					Message: "Module source \"git::https://hashicorp.com/consul.git\" is not pinned",
					Range: hcl.Range{
						Filename: "module.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 51},
					},
				},
			},
		},
		{
			Name: "generic git (git::ssh) module reference is not pinned",
			Content: `
module "unpinned_generic_git_ssh" {
  source = "git::ssh://git@github.com/owner/repo.git"
}
`,
			Expected: tflint.Issues{
				{
					Rule:    NewTerraformModulePinnedSourceRule(),
					Message: "Module source \"git::ssh://git@github.com/owner/repo.git\" is not pinned",
					Range: hcl.Range{
						Filename: "module.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 54},
					},
				},
			},
		},
		{
			Name: "generic git (git::https) module reference is default",
			Content: `
module "default_generic_git_https" {
  source = "git::https://hashicorp.com/consul.git?ref=master"
}
`,
			Expected: tflint.Issues{
				{
					Rule:    NewTerraformModulePinnedSourceRule(),
					Message: "Module source \"git::https://hashicorp.com/consul.git?ref=master\" uses default ref \"master\"",
					Range: hcl.Range{
						Filename: "module.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 62},
					},
				},
			},
		},
		{
			Name: "generic git (git::ssh) module reference is default",
			Content: `
module "default_generic_git_ssh" {
  source = "git::ssh://git@github.com/owner/repo.git?ref=master"
}
`,
			Expected: tflint.Issues{
				{
					Rule:    NewTerraformModulePinnedSourceRule(),
					Message: "Module source \"git::ssh://git@github.com/owner/repo.git?ref=master\" uses default ref \"master\"",
					Range: hcl.Range{
						Filename: "module.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 65},
					},
				},
			},
		},
		{
			Name: "generic git (git::https) module reference is pinned",
			Content: `
module "pinned_generic_git_https" {
  source = "git::https://hashicorp.com/consul.git?ref=pinned"
}
`,
			Expected: tflint.Issues{},
		},
		{
			Name: "generic git (git::ssh) module reference is pinned",
			Content: `
module "pinned_generic_git_ssh" {
  source = "git::ssh://git@github.com/owner/repo.git?ref=pinned"
}
`,
			Expected: tflint.Issues{},
		},
		{
			Name: "mercurial module is not pinned",
			Content: `
module "default_mercurial" {
  source = "hg::http://hashicorp.com/consul.hg"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewTerraformModulePinnedSourceRule(),
					Message: "Module source \"hg::http://hashicorp.com/consul.hg\" is not pinned",
					Range: hcl.Range{
						Filename: "module.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 48},
					},
				},
			},
		},
		{
			Name: "mercurial module reference is default",
			Content: `
module "default_mercurial" {
  source = "hg::http://hashicorp.com/consul.hg?rev=default"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewTerraformModulePinnedSourceRule(),
					Message: "Module source \"hg::http://hashicorp.com/consul.hg?rev=default\" uses default rev \"default\"",
					Range: hcl.Range{
						Filename: "module.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 60},
					},
				},
			},
		},
		{
			Name: "mercurial module reference is pinned",
			Content: `
module "pinned_mercurial" {
  source = "hg::http://hashicorp.com/consul.hg?rev=pinned"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewTerraformModulePinnedSourceRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"module.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
