package odl

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccVInterface_Basic(t *testing.T) {
	tenantName := "vtn3"
	bridgeName := "vbr3"
	interfaceName := "itr1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVInterfaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckVInterfaceConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVInterfaceExists("odl_vinterface.vinterface1"),
					resource.TestCheckResourceAttr(
						"odl_vinterface.vinterface1", "tenant_name", tenantName),
					resource.TestCheckResourceAttr(
						"odl_vinterface.vinterface1", "bridge_name", bridgeName),
					resource.TestCheckResourceAttr(
						"odl_vinterface.vinterface1", "interface_name", interfaceName),
				),
			},
		},
	})
}

func testAccCheckVInterfaceDestroy(s *terraform.State) error {

	rs, ok := s.RootModule().Resources["odl_VInterface.VInterface1"]

	if !ok {
		return fmt.Errorf("Not found: odl_VInterface.VInterface1")
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
	present, err := CheckResponseVInterfaceExists(response, tenantName, bridgeName, interfaceName)
	if err != nil {
		log.Println("[ERROR] VInterface Read failed")
		return fmt.Errorf("[ERROR] VInterface could not be read %v", err)
	}
	if present {
		log.Println("[INFO] VInterface with name " + interfaceName + " found")
		return fmt.Errorf("[ERROR] VInterface with name " + interfaceName + "was found")
	}
	return nil
}

func testAccCheckVInterfaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VInterface ID is set")
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
		present, err := CheckResponseVInterfaceExists(response, tenantName, bridgeName, interfaceName)
		if err != nil {
			log.Println("[ERROR] VInterface Read failed")
			return fmt.Errorf("[ERROR] VInterface could not be read %v", err)
		}
		if !present {
			log.Println("[INFO] VInterface with name " + interfaceName + "was not found")
			return fmt.Errorf("[ERROR] VInterface with name " + interfaceName + "was not found")
		}
		return nil
	}
}

const testAccCheckVInterfaceConfigBasic = `
resource "odl_vinterface" "vinterface1"{
     	tenant_name = "vtn3"
		bridge_name = "vbr3"
		interface_name = "itr1"
}`
