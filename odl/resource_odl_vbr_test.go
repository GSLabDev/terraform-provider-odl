package odl

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccVbr_Basic(t *testing.T) {
	tenantName := "vtn1"
	bridgeName := "vbr1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVbrDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckVbrConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVbrExists("odl_vbr.vbr1"),
					resource.TestCheckResourceAttr(
						"odl_vbr.vbr1", "tenant_name", tenantName),
					resource.TestCheckResourceAttr(
						"odl_vbr.vbr1", "bridge_name", bridgeName),
				),
			},
		},
	})
}

func testAccCheckVbrDestroy(s *terraform.State) error {

	rs, ok := s.RootModule().Resources["odl_vbr.vbr1"]

	if !ok {
		return fmt.Errorf("Not found: odl_vbr.vbr1")
	}

	tenantName := rs.Primary.Attributes["tenant_name"]
	bridgeName := rs.Primary.Attributes["bridge_name"]
	config := testAccProvider.Meta().(*Config)

	response, err := config.GetRequest("restconf/operational/vtn:vtns")
	if err != nil {
		log.Printf("[ERROR] POST Request failed")
		return err
	}
	present, err := CheckResponseVbrExists(response, tenantName, bridgeName)
	if err != nil {
		log.Println("[ERROR] Vbr Read failed")
		return fmt.Errorf("[ERROR] Vbr could not be read %v", err)
	}
	if present {
		log.Println("[INFO] Vbr with name " + bridgeName + " found")
		return fmt.Errorf("[ERROR] Vbr with name " + bridgeName + "was found")
	}
	return nil
}

func testAccCheckVbrExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Vbr ID is set")
		}
		tenantName := rs.Primary.Attributes["tenant_name"]
		bridgeName := rs.Primary.Attributes["bridge_name"]

		config := testAccProvider.Meta().(*Config)

		response, err := config.GetRequest("restconf/operational/vtn:vtns")
		if err != nil {
			log.Printf("[ERROR] POST Request failed")
			return err
		}
		present, err := CheckResponseVbrExists(response, tenantName, bridgeName)
		if err != nil {
			log.Println("[ERROR] Vbr Read failed")
			return fmt.Errorf("[ERROR] Vbr could not be read %v", err)
		}
		if !present {
			log.Println("[INFO] Vbr with name " + bridgeName + "was not found")
			return fmt.Errorf("[ERROR] Vbr with name " + bridgeName + "was not found")
		}
		return nil
	}
}

const testAccCheckVbrConfigBasic = `
resource "odl_vbr" "vbr1"{
     	tenant_name = "vtn1"
		bridge_name = "vbr1"
}`
