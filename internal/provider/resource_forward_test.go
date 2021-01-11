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

func testAccResourceForwardConfig(domain string, host string, url string) string {
	return fmt.Sprintf(`
data "domeneshop_domains" "test" {
  domain = "%s"
}

resource "domeneshop_forward" "test" {
  domain_id = data.domeneshop_domains.test.domains.0.id
  host      = "%s"
  url       = "%s"
}
`, domain, host, url)
}
