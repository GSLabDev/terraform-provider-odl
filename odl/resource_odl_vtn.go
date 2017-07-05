package odl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOdlVtn() *schema.Resource {
	return &schema.Resource{
		Create: resourceVtnAdd,
		Read:   resourceVtnRead,
		Delete: resourceVtnDelete,
		Schema: map[string]*schema.Schema{
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"operation": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateOperation,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"hard_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}
func resourceVtnAdd(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	tenantName := d.Get("tenant_name").(string)

	var body map[string]interface{}
	var input map[string]string
	input = make(map[string]string)
	log.Println("[INFO] Creating Vtn with name " + tenantName)
	input["tenant-name"] = tenantName
	input["update-mode"] = "UPDATE"
	if operation, found := d.GetOk("operation"); found {
		input["operation"] = operation.(string)
	}
	if description, found := d.GetOk("description"); found {
		input["description"] = description.(string)
	}
	if idleTimeout, found := d.GetOk("idle_timeout"); found {
		input["idle-timeout"] = string(idleTimeout.(int))
	}
	if hardTimeout, found := d.GetOk("hard_timeout"); found {
		input["hard-timeout"] = string(hardTimeout.(int))
	}
	log.Println("[INFO] All options collected for Vtn with name " + tenantName)
	body = make(map[string]interface{})
	body["input"] = input
	response, err := config.PostRequest("restconf/operations/vtn:update-vtn", body)
	if err != nil {
		log.Printf("[ERROR] POST Request failed")
		return err
	}
	isCreated, output, errorOutput, err := Status(response)
	if isCreated {
		d.SetId(tenantName + output.Output.Status)
	} else {
		if errorOutput != nil {
			log.Printf("[ERROR] While creating vtn %s", errorOutput.Errors.Error[0].Message)
			return fmt.Errorf("[ERROR] While creating vtn %s", errorOutput.Errors.Error[0].Message)
		}
		if err != nil {
			return fmt.Errorf("[ERROR] Whlie creating vtn %s", err.Error())
		}
	}

	return nil
}
func resourceVtnRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	tenantName := d.Get("tenant_name").(string)
	log.Println("[INFO] Read Vtn with name " + tenantName)
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
		log.Println("[INFO] Vtn with name " + tenantName + "was not found")
		d.SetId("")
	}
	return nil
}
func resourceVtnDelete(d *schema.ResourceData, meta interface{}) error {
	err := resourceVtnRead(d, meta)
	if d.Id() == "" {
		return fmt.Errorf("[ERROR] vtn does not exists")
	}
	config := meta.(*Config)
	tenantName := d.Get("tenant_name").(string)

	var body map[string]interface{}
	var input map[string]string
	input = make(map[string]string)

	input["tenant-name"] = tenantName
	body = make(map[string]interface{})
	body["input"] = input
	log.Println("[INFO] All options collected for Vtn with name " + tenantName)
	log.Println("[INFO] Preparing to destroy Vtn with name " + tenantName)

	response, err := config.PostRequest("restconf/operations/vtn:remove-vtn", body)
	if err != nil {
		log.Printf("[ERROR] POST Request failed")
		return err
	}
	isDestroyed, _, errorOutput, err := Status(response)
	if isDestroyed {
		d.SetId("")
	} else {
		if errorOutput != nil {
			log.Printf("[ERROR] While destroying vtn %s", errorOutput.Errors.Error[0].Message)
			return fmt.Errorf("[ERROR] While creating vtn %s", errorOutput.Errors.Error[0].Message)
		}
		if err != nil {
			return fmt.Errorf("[ERROR] Whlie creating vtn %s", err.Error())
		}
	}

	return nil
}
