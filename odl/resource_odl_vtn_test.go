package odl

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccVtn_Basic(t *testing.T) {
	tenantName := "vtn2"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVtnDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckVtnConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVtnExists("odl_vtn.vtn2"),
					resource.TestCheckResourceAttr(
						"odl_vtn.vtn2", "tenant_name", tenantName),
				),
			},
		},
	})
}

func testAccCheckVtnDestroy(s *terraform.State) error {

	rs, ok := s.RootModule().Resources["odl_vtn.vtn2"]

	if !ok {
		return fmt.Errorf("Not found: odl_vtn.vtn2")
	}

	tenantName := rs.Primary.Attributes["tenant_name"]
	config := testAccProvider.Meta().(*Config)

	response, err := config.GetRequest("restconf/operational/vtn:vtns")
	if err != nil {
		log.Printf("[ERROR] POST Request failed")
		return err
	}
	present, err := CheckResponseVtnExists(response, tenantName)
	if err != nil {
		log.Println("[ERROR] Vtn Read failed")
		return fmt.Errorf("[ERROR] Vtn could not be read %v", err)
	}
	if present {
		log.Println("[INFO] Vtn with name " + tenantName + " found")
		return fmt.Errorf("[ERROR] Vtn with name " + tenantName + "was found")
	}
	return nil
}

func testAccCheckVtnExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Vtn ID is set")
		}
		tenantName := rs.Primary.Attributes["tenant_name"]
		config := testAccProvider.Meta().(*Config)

		response, err := config.GetRequest("restconf/operational/vtn:vtns")
		if err != nil {
			log.Printf("[ERROR] POST Request failed")
			return err
		}
		present, err := CheckResponseVtnExists(response, tenantName)
		if err != nil {
			log.Println("[ERROR] Vtn Read failed")
			return fmt.Errorf("[ERROR] Vtn could not be read %v", err)
		}
		if !present {
			log.Println("[INFO] Vtn with name " + tenantName + " found")
			return fmt.Errorf("[ERROR] Vtn with name " + tenantName + "was found")
		}
		return nil
	}
}

const testAccCheckVtnConfigBasic = `
resource "odl_vtn" "vtn2"{
     	tenant_name = "vtn2"
}`
