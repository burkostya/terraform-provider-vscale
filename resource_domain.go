package main

import (
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	vscale "github.com/vganyn/go-vscale"
)

func resourceDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceDomainCreate,
		Read:   resourceDomainRead,
		Exists: resourceDomainExists,
		Delete: resourceDomainDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDomainCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*vscale.WebClient)

	name := d.Get("name").(string)

	domain, _, err := client.Domain.Create(name)
	if err != nil {
		return errors.Wrap(err, "creating domain failed")
	}

	d.SetId(strconv.FormatInt(domain.ID, 10))

	return nil
}

func resourceDomainRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*vscale.WebClient)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	domain, _, err := client.Domain.Get(id)
	if err != nil {
		return errors.Wrap(err, "getting domain failed")
	}

	d.Set("name", domain.Name)

	return nil
}

func resourceDomainExists(d *schema.ResourceData, m interface{}) (bool, error) {
	client := m.(*vscale.WebClient)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return true, errors.Wrap(err, "invalid id")
	}

	_, res, err := client.Domain.Get(id)
	if res.StatusCode == http.StatusNotFound {
		return false, nil
	}
	if err != nil {
		return true, errors.Wrap(err, "getting domain failed")
	}

	return true, nil
}

func resourceDomainDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*vscale.WebClient)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	_, res, err := client.Domain.Remove(id)
	if res.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil {
		return errors.Wrap(err, "removing domain failed")
	}

	return nil
}
