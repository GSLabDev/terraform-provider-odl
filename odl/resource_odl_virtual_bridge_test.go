package odl

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccOdlVirtualBridge_Basic(t *testing.T) {
	tenantName := "terraformVtn"
	bridgeName := "terraformBridge"
	resourceName := "odl_virtual_bridge.firstVirtualBridge"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVirtualBridgeDestroy(resourceName),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckVirtualBridgeConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVirtualBridgeExists(resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "tenant_name", tenantName),
					resource.TestCheckResourceAttr(
						resourceName, "bridge_name", bridgeName),
				),
			},
		},
	})
}

func testAccCheckVirtualBridgeDestroy(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: " + n)
		}

		tenantName := rs.Primary.Attributes["tenant_name"]
		bridgeName := rs.Primary.Attributes["bridge_name"]
		config := testAccProvider.Meta().(*Config)

		response, err := config.GetRequest("restconf/operational/vtn:vtns")
		if err != nil {
			log.Printf("[ERROR] POST Request failed")
			return err
		}
		present, err := CheckResponseVirtualBridgeExists(response, tenantName, bridgeName)
		if err != nil {
			log.Println("[ERROR] Virtual Bridge Read failed")
			return fmt.Errorf("[ERROR] Virtual Bridge could not be read %v", err)
		}
		if present {
			log.Println("[DEBUG] Virtual Bridge with name " + bridgeName + " found")
			return fmt.Errorf("[ERROR] Virtual Bridge with name " + bridgeName + "was found")
		}
		return nil
	}
}

func testAccCheckVirtualBridgeExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Virtual Bridge ID is set")
		}
		tenantName := rs.Primary.Attributes["tenant_name"]
		bridgeName := rs.Primary.Attributes["bridge_name"]

		config := testAccProvider.Meta().(*Config)

		response, err := config.GetRequest("restconf/operational/vtn:vtns")
		if err != nil {
			log.Printf("[ERROR] POST Request failed")
			return err
		}
		present, err := CheckResponseVirtualBridgeExists(response, tenantName, bridgeName)
		if err != nil {
			log.Println("[ERROR] Virtual Bridge Read failed")
			return fmt.Errorf("[ERROR] Virtual Bridge could not be read %v", err)
		}
		if !present {
			log.Println("[DEBUG] Virtual Bridge with name " + bridgeName + "was not found")
			return fmt.Errorf("[ERROR] Virtual Bridge with name " + bridgeName + "was not found")
		}
		return nil
	}
}

const testAccCheckVirtualBridgeConfigBasic = `
resource "odl_virtual_tenant_network" "firstVtn" {
  tenant_name  = "terraformVtn"
  operation    = "ADD"
  description  = "operation can be ADD or SET only"
  idle_timeout = 56
  hard_timeout = 58
}
  
resource "odl_virtual_bridge" "firstVirtualBridge" {
  tenant_name  = "${odl_virtual_tenant_network.firstVtn.tenant_name}"
  bridge_name  = "terraformBridge"
  operation    = "SET"
  description  = "operation can be ADD or SET only"
  age_interval = 577
}`
