package odl

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOdlVirtualBridge() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualBridgeAdd,
		Read:   resourceVirtualBridgeRead,
		Delete: resourceVirtualBridgeDelete,
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
			"age_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"bridge_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}
func resourceVirtualBridgeAdd(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	tenantName := d.Get("tenant_name").(string)
	bridgeName := d.Get("bridge_name").(string)

	var body map[string]interface{}
	var input map[string]string
	input = make(map[string]string)

	log.Println("[DEBUG] Creating Virtual Bridge with name " + bridgeName)
	input["tenant-name"] = tenantName
	input["update-mode"] = "UPDATE"
	input["bridge-name"] = bridgeName
	if operation, found := d.GetOk("operation"); found {
		input["operation"] = operation.(string)
	}
	if description, found := d.GetOk("description"); found {
		input["description"] = description.(string)
	}
	if idleTimeout, found := d.GetOk("age_interval"); found {
		input["age-interval"] = strconv.Itoa(idleTimeout.(int))
	}
	log.Println("[DEBUG] All options collected for Virtual Bridge with name " + tenantName)
	body = make(map[string]interface{})
	body["input"] = input
	response, err := config.PostRequest("restconf/operations/vtn-vbridge:update-vbridge", body)
	if err != nil {
		log.Printf("[ERROR] POST Request failed")
		return err
	}
	isCreated, output, errorOutput, err := Status(response)
	if isCreated {
		d.SetId(tenantName + bridgeName + output.Output.Status)
	} else {
		if errorOutput != nil {
			log.Printf("[ERROR] While creating vbr %s", errorOutput.Errors.Error[0].Message)
			return fmt.Errorf("[ERROR] While creating vbr %s", errorOutput.Errors.Error[0].Message)
		}
		if err != nil {
			return fmt.Errorf("[ERROR] Whlie creating vbr %s", err.Error())
		}
	}

	return nil
}
func resourceVirtualBridgeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	tenantName := d.Get("tenant_name").(string)
	bridgeName := d.Get("bridge_name").(string)

	log.Println("[DEBUG] Read Bridge with name " + bridgeName)
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
		log.Println("[DEBUG] VirtualBridge with name " + bridgeName + "was not found")
		d.SetId("")
	}
	return nil
}
func resourceVirtualBridgeDelete(d *schema.ResourceData, meta interface{}) error {
	err := resourceVirtualBridgeRead(d, meta)
	if d.Id() == "" {
		return fmt.Errorf("[ERROR] vbr does not exists")
	}
	config := meta.(*Config)
	tenantName := d.Get("tenant_name").(string)
	bridgeName := d.Get("bridge_name").(string)
	var body map[string]interface{}
	var input map[string]string
	input = make(map[string]string)

	input["tenant-name"] = tenantName
	input["bridge-name"] = bridgeName
	body = make(map[string]interface{})
	body["input"] = input

	log.Println("[DEBUG] All options collected for Virtual Bridge with name " + bridgeName)
	log.Println("[DEBUG] Preparing to destroy Virtual Bridge with name " + bridgeName)

	response, err := config.PostRequest("restconf/operations/vtn-vbridge:remove-vbridge", body)
	if err != nil {
		log.Printf("[ERROR] POST Request failed")
		return err
	}
	isDestroyed, _, errorOutput, err := Status(response)
	if isDestroyed {
		d.SetId("")
	} else {
		if errorOutput != nil {
			log.Printf("[ERROR] While destroying vbr %s", errorOutput.Errors.Error[0].Message)
			return fmt.Errorf("[ERROR] While creating vbr %s", errorOutput.Errors.Error[0].Message)
		}
		if err != nil {
			return fmt.Errorf("[ERROR] Whlie destroying vbr %s", err.Error())
		}
	}

	return nil
}
