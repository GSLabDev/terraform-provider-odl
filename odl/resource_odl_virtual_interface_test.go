package odl

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccOdlVirtualInterface_Basic(t *testing.T) {
	tenantName := "terraformVtn"
	bridgeName := "terraformBridge"
	interfaceName := "terraformInterface"
	resourceName := "odl_virtual_interface.firstInterface"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVirtualInterfaceDestroy(resourceName),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckVirtualInterfaceConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVirtualInterfaceExists(resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "tenant_name", tenantName),
					resource.TestCheckResourceAttr(
						resourceName, "bridge_name", bridgeName),
					resource.TestCheckResourceAttr(
						resourceName, "interface_name", interfaceName),
				),
			},
		},
	})
}

func testAccCheckVirtualInterfaceDestroy(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: " + n)
		}

		tenantName := rs.Primary.Attributes["tenant_name"]
		bridgeName := rs.Primary.Attributes["bridge_name"]
		interfaceName := rs.Primary.Attributes["interface_name"]
		config := testAccProvider.Meta().(*Config)

		response, err := config.GetRequest("restconf/operational/vtn:vtns")
		if err != nil {
			log.Printf("[ERROR] POST Request failed")
			return err
		}
		present, err := CheckResponseVirtualInterfaceExists(response, tenantName, bridgeName, interfaceName)
		if err != nil {
			log.Println("[ERROR] Virtual Interface Read failed")
			return fmt.Errorf("[ERROR] Virtual Interface could not be read %v", err)
		}
		if present {
			log.Println("[DEBUG] Virtual Interface with name " + interfaceName + " found")
			return fmt.Errorf("[ERROR] Virtual Interface with name " + interfaceName + "was found")
		}
		return nil
	}
}

func testAccCheckVirtualInterfaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Virtual Interface ID is set")
		}
		tenantName := rs.Primary.Attributes["tenant_name"]
		bridgeName := rs.Primary.Attributes["bridge_name"]
		interfaceName := rs.Primary.Attributes["interface_name"]

		config := testAccProvider.Meta().(*Config)

		response, err := config.GetRequest("restconf/operational/vtn:vtns")
		if err != nil {
			log.Printf("[ERROR] POST Request failed")
			return err
		}
		present, err := CheckResponseVirtualInterfaceExists(response, tenantName, bridgeName, interfaceName)
		if err != nil {
			log.Println("[ERROR] Virtual Interface Read failed")
			return fmt.Errorf("[ERROR] Virtual Interface could not be read %v", err)
		}
		if !present {
			log.Println("[DEBUG] Virtual Interface with name " + interfaceName + "was not found")
			return fmt.Errorf("[ERROR] Virtual Interface with name " + interfaceName + "was not found")
		}
		return nil
	}
}

const testAccCheckVirtualInterfaceConfigBasic = `
resource "odl_virtual_tenant_network" "firstVtn" {
  tenant_name  = "terraformVtn"
  operation    = "ADD"
  description  = "operation can be ADD or SET only"
  idle_timeout = 56
  hard_timeout = 58
}
	
resource "odl_virtual_bridge" "firstVbr" {
  tenant_name  = "${odl_virtual_tenant_network.firstVtn.tenant_name}"
  bridge_name  = "terraformBridge"
  operation    = "SET"
  description  = "operation can be ADD or SET only"
  age_interval = 577
}

resource "odl_virtual_interface" "firstInterface" {
  tenant_name    = "${odl_virtual_tenant_network.firstVtn.tenant_name}"
  bridge_name    = "${odl_virtual_bridge.firstVbr.bridge_name}"
  description    = "operation can be ADD or SET only"
  interface_name = "terraformInterface"
  enabled        = true
  terminal_name  = "ter1"
}`
