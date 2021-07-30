package neptune

// ATTENTION!
// This rule was autogenerated!
// Before making changes, consider updating the generator.

import (
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/hclcontext"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/provider"
	"github.com/aquasecurity/tfsec/pkg/result"
	"github.com/aquasecurity/tfsec/pkg/rule"
	"github.com/aquasecurity/tfsec/pkg/severity"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		Provider:       provider.AWSProvider,
		Service:   "neptune",
		ShortCode: "enable-storage-encryption",
		Documentation: rule.RuleDocumentation{
			Summary:     "Neptune storage must be encrypted at rest",
			Explanation: `Encryption of Neptune storage ensures that if their is compromise of the disks, the data is still protected.`,
			Impact:      "Unencrypted sensitive data is vulnerable to compromise.",
			Resolution:  "Enable encryption of Neptune storage",
			BadExample: []string{  `
resource "aws_neptune_cluster" "bad_example" {
  cluster_identifier                  = "neptune-cluster-demo"
  engine                              = "neptune"
  backup_retention_period             = 5
  preferred_backup_window             = "07:00-09:00"
  skip_final_snapshot                 = true
  iam_database_authentication_enabled = true
  apply_immediately                   = true
  storage_encrypted = false
}
`},
			GoodExample: []string{ `
resource "aws_neptune_cluster" "good_example" {
  cluster_identifier                  = "neptune-cluster-demo"
  engine                              = "neptune"
  backup_retention_period             = 5
  preferred_backup_window             = "07:00-09:00"
  skip_final_snapshot                 = true
  iam_database_authentication_enabled = true
  apply_immediately                   = true
  storage_encrypted = true
}
`},
			Links: []string{
				"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/neptune_cluster#storage_encrypted",
			},
		},
		RequiredTypes:  []string{ 
			"resource",
		},
		RequiredLabels: []string{ 
			"aws_neptune_cluster",
		},
		DefaultSeverity: severity.High, 
		CheckFunc: func(set result.Set, resourceBlock block.Block, _ *hclcontext.Context){
			if storageEncryptedAttr := resourceBlock.GetAttribute("storage_encrypted"); storageEncryptedAttr.IsNil() { // alert on use of default value
				set.AddResult().
					WithDescription("Resource '%s' uses default value for storage_encrypted", resourceBlock.FullName())
			} else if storageEncryptedAttr.IsFalse() {
				set.AddResult().
					WithDescription("Resource '%s' does not have storage_encrypted set to true", resourceBlock.FullName()).
					WithAttribute(storageEncryptedAttr)
			}
		},
	})
}
