package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	vscale "github.com/vscale/go-vscale"
)

// Provider returns a schema.Provider for VScale.
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"vscale_scalet":  resourceScalet(),
			"vscale_ssh_key": resourceSSHKey(),
		},
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSCALE_API_TOKEN", nil),
				Description: "The token key for API operations.",
			},
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := vscale.NewClient(d.Get("token").(string))

	return client, nil
}
