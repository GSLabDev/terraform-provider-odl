package odl

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccOdlVirtualTenantNetwork_Basic(t *testing.T) {
	resourceName := "odl_virtual_tenant_network.vtn2"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVirtualTenantNetworkDestroy(resourceName),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckVirtualTenantNetworkConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVirtualTenantNetworkExists(resourceName),
					resource.TestCheckResourceAttr(
						"odl_virtual_tenant_network.vtn2", "tenant_name", "terraformVirtualTenantNetwork"),
				),
			},
		},
	})
}

func testAccCheckVirtualTenantNetworkDestroy(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: odl_virtual_tenant_network.vtn2")
		}

		tenantName := rs.Primary.Attributes["tenant_name"]
		config := testAccProvider.Meta().(*Config)

		response, err := config.GetRequest("restconf/operational/vtn:vtns")
		if err != nil {
			log.Printf("[ERROR] POST Request failed")
			return err
		}
		present, err := CheckResponseVirtualTenantNetworkExists(response, tenantName)
		if err != nil {
			log.Println("[ERROR] VirtualTenantNetwork Read failed")
			return fmt.Errorf("[ERROR] VirtualTenantNetwork could not be read %v", err)
		}
		if present {
			log.Println("[DEBUG] VirtualTenantNetwork with name " + tenantName + " found")
			return fmt.Errorf("[ERROR] VirtualTenantNetwork with name " + tenantName + "was found")
		}
		return nil
	}
}

func testAccCheckVirtualTenantNetworkExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VirtualTenantNetwork ID is set")
		}
		tenantName := rs.Primary.Attributes["tenant_name"]
		config := testAccProvider.Meta().(*Config)

		response, err := config.GetRequest("restconf/operational/vtn:vtns")
		if err != nil {
			log.Printf("[ERROR] POST Request failed")
			return err
		}
		present, err := CheckResponseVirtualTenantNetworkExists(response, tenantName)
		if err != nil {
			log.Println("[ERROR] VirtualTenantNetwork Read failed")
			return fmt.Errorf("[ERROR] VirtualTenantNetwork could not be read %v", err)
		}
		if !present {
			log.Println("[DEBUG] VirtualTenantNetwork with name " + tenantName + " found")
			return fmt.Errorf("[ERROR] VirtualTenantNetwork with name " + tenantName + "was found")
		}
		return nil
	}
}

const testAccCheckVirtualTenantNetworkConfigBasic = `
resource "odl_virtual_tenant_network" "vtn2" {
  tenant_name = "terraformVirtualTenantNetwork"
}`
