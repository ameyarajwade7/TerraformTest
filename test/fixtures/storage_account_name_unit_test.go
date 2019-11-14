package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestUT_StorageAccountName(t *testing.T) {
	t.Parallel()

	// Test cases for storage account name conversion logic
	testCases := map[string]string{
		"TestWebsiteName": "testwebsitenamedata001",
		// "ALLCAPS":         "allcapsdata001",
		// "S_p-e(c)i.a_l":   "specialdata001",
		// "A1phaNum321":     "a1phanum321data001",
		// "E5e-y7h_ng":      "e5ey7hngdata001",
	}

	for input := range testCases {
	
		tfOptions := &terraform.Options{
			TerraformDir: "storage-account-name",
			Vars: map[string]interface{}{
				"website_name": input,
			},
		}
		tfOptionsa := &terraform.Options{
			TerraformDir: "storage-account-name",
		}

		// Terraform init and plan only
		tfPlanOutput := "terraform.tfplan"
		terraform.Init(t, tfOptions)
		terraform.RunTerraformCommand(t, tfOptions, terraform.FormatArgs(tfOptions, "plan", "-out="+tfPlanOutput)...)
		// Read and parse the plan output
		terraform.RunTerraformCommandAndGetStdoutE(
			t, tfOptions, terraform.FormatArgs(tfOptionsa, "show", "-json", tfPlanOutput)...,
		)

		// Validate the test result
		// for _, mod := range planJSON. {
		// 	if len(mod.Path) == 2 && mod.Path[0] == "root" && mod.Path[1] == "staticwebpage" {
		// 		actual := mod.Resources["azurerm_storage_account.main"].Attributes["name"].New
		// 		if actual != expected {
		// 			t.Fatalf("Expect %v, but found %v", expected, actual)
		// 		}
		// 	}
		// }
	}
}
