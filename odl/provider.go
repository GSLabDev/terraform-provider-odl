package odl

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{

			"server_ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "odl server ip",
				DefaultFunc: schema.EnvDefaultFunc("ODL_SERVER_IP", nil),
			},

			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "odl server port",
				DefaultFunc: schema.EnvDefaultFunc("ODL_SERVER_PORT", nil),
			},

			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "odl user name",
				DefaultFunc: schema.EnvDefaultFunc("ODL_SERVER_USER", nil),
			},

			"user_password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "odl user password",
				DefaultFunc: schema.EnvDefaultFunc("ODL_SERVER_PASSWORD", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"odl_vtn": resourceOdlVtn(),
			"odl_vbr": resourceOdlVbr(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := Config{
		ServerIP: d.Get("server_ip").(string),
		Port:     d.Get("port").(int),
		Username: d.Get("user_name").(string),
		Password: d.Get("user_password").(string),
	}
	return config.checkConnection()
}
