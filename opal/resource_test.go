package opal

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/opalsecurity/opal-go"
)

var knownCustomAppID = os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ID")
var knownCustomAppAdminOwnerID = os.Getenv("OPAL_TEST_KNOWN_CUSTOM_APP_ADMIN_OWNER_ID")
var knownRequestTemplateID = os.Getenv("OPAL_TEST_KNOWN_REQUEST_TEMPLATE_ID")

func TestAccResource_CRUD(t *testing.T) {
	baseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_resource." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResource(baseName, baseName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),                             // Verify that the name was set.
					resource.TestCheckResourceAttr(resourceName, "description", ""),                            // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "max_duration", "0"),                          // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "admin_owner_id", knownCustomAppAdminOwnerID), // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "require_manager_approval", "false"),          // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "require_support_ticket", "false"),            // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "require_mfa_to_approve", "false"),            // Verify that optional works.
					resource.TestCheckResourceAttr(resourceName, "auto_approval", "false"),                     // Verify that optional works.
				),
			},
			{
				Config: testAccResourceResource(baseName, baseName+"_changed", `description = "test desc"`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName+"_changed"), // Verify that updating the name works.
					resource.TestCheckResourceAttr(resourceName, "description", "test desc"),  // Verify that updating the description works.
				),
			},
		},
	})
}

// TestAccResource_SetOnCreate tests that setting attributes on creation
// works.
func TestAccResource_SetOnCreate(t *testing.T) {
	baseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_resource." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResource(baseName, baseName, fmt.Sprintf(`
description = "test desc"
require_manager_approval = true
require_support_ticket = true
max_duration = 30
request_template_id = "%s"
`, knownRequestTemplateID)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", baseName),
					resource.TestCheckResourceAttr(resourceName, "description", "test desc"),
					resource.TestCheckResourceAttr(resourceName, "auto_approval", "false"),
				),
			},
		},
	})
}

// TestAccResource_SetOnCreate tests that setting auto approve on creation works.
func TestAccResource_SetOnCreate_AutoApproval(t *testing.T) {
	baseName := "tf_acc_test_resource_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resourceName := "opal_resource." + baseName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceResource(baseName, baseName, `
auto_approval = true
`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "auto_approval", "true"),
				),
			},
		},
	})
}

// XXX: Test combination plan with owner and resource.
// XXX: Test metadata / Remote ID

func testAccResourceResource(tfName, name, additional string) string {
	return fmt.Sprintf(`
resource "opal_resource" "%s" {
	name = "%s"
	app_id = "%s"
	resource_type = "CUSTOM"

	%s
}
`, tfName, name, knownCustomAppID, additional)
}

func testAccCheckResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*opal.APIClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opal_resource" {
			continue
		}
		resource, _, err := client.ResourcesApi.GetResource(context.Background(), rs.Primary.ID).Execute()
		if err == nil {
			if resource != nil {
				return fmt.Errorf("Opal resource still exists: %s", rs.Primary.ID)
			}
			return nil
		}
		if !strings.Contains(err.Error(), "404 Not Found") {
			return err
		}
	}

	return nil
}
