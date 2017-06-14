package odl

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	//"log"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{

			"ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "domain",
				DefaultFunc: schema.EnvDefaultFunc("ODL_IP", nil),
			},

			"port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ip",
				DefaultFunc: schema.EnvDefaultFunc("ODL_PORT", nil),
			},

			"user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "user",
				DefaultFunc: schema.EnvDefaultFunc("ODL_USER", nil),
			},

			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "password",
				DefaultFunc: schema.EnvDefaultFunc("ODL_PASSWORD", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"odl_networkTopology": networkTopology(),
		},

		ConfigureFunc: providerConfigure,
	}
}

/*func providerConfigure(d *schema.ResourceData) (interface{}, error) {

 IP:=	d.Get("ip").(string),
 Port:=	d.Get("port").(string),
 Username:=	d.Get("user").(string),
 Password:=	d.Get("password").(string),

 return SimpleFakeApi.New(IP, Port,Username,Password )

}*/
