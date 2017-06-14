package odl

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"net/http"
)

func networkTopology() *schema.Resource {
	return &schema.Resource{
		Create: networkTopologyAdd,
		Read:   networkTopologyRead,
		Delete: networkTopologyDelete,
		Update: networkTopologyUpdate,
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resturl": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func networkTopologyRead(d *schema.ResourceData, meta interface{}) error {

	//Baseurl for the setup
	//baseurl := "http://192.168.56.102:8181"	
	baseurl := d.Get("user").(string) + ":" + d.Get("password").(string) + "@" + d.Get("ip").(string) + ":" + d.Get("port").(string) + "/"
	resturl := d.Get("resturl").(string)
	
	
	//concatinate the resturl to baseurl for specific API call
	url := baseurl + resturl
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request to the server")
		return nil
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	return nil
}

func networkTopologyAdd(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func networkTopologyDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func networkTopologyUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
