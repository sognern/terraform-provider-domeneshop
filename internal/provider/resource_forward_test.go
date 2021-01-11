package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceForward(t *testing.T) {
	domain := os.Getenv("DOMENESHOP_DOMAIN")
	host := acctest.RandString(6)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// TODO: CheckDestroy: ,
		Steps: []resource.TestStep{
			{
				// test create
				Config: testAccResourceForwardConfig(domain, host, "https://example.com/foo"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("domeneshop_forward.test", "host", host),
					resource.TestCheckResourceAttr("domeneshop_forward.test", "url", "https://example.com/foo"),
					resource.TestCheckResourceAttr("domeneshop_forward.test", "frame", "false"),
				),
			},
			{
				// test update
				Config: testAccResourceForwardConfig(domain, host, "https://example.com/bar"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("domeneshop_forward.test", "host", host),
					resource.TestCheckResourceAttr("domeneshop_forward.test", "url", "https://example.com/bar"),
					resource.TestCheckResourceAttr("domeneshop_forward.test", "frame", "false"),
				),
			},
		},
	})
}

func TestAccResourceForward_Import(t *testing.T) {
	domainID := os.Getenv("DOMENESHOP_DOMAIN_ID")
	if domainID == "" {
		t.Skip(`Skipping test because "DOMENESHOP_DOMAIN_ID" is not set`)
	}
	host := acctest.RandString(6)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// TODO: CheckDestroy: ,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceForwardImportConfig(domainID, host),
			},
			{
				ResourceName:        "domeneshop_forward.test",
				ImportStateIdPrefix: fmt.Sprintf("%s/", domainID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccResourceForwardConfig(domain string, host string, url string) string {
	return fmt.Sprintf(`
data "domeneshop_domains" "test" {
  domain = "%s"
}

resource "domeneshop_forward" "test" {
  domain_id = data.domeneshop_domains.test.domains[0].id
  host      = "%s"
  url       = "%s"
}
`, domain, host, url)
}

func testAccResourceForwardImportConfig(domain string, host string) string {
	return fmt.Sprintf(`
resource "domeneshop_forward" "test" {
  domain_id = %s
  host      = "%s"
  url       = "https://example.com/"
}
`, domain, host)
}
