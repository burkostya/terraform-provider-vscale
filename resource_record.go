package main

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	vscale "github.com/vganyn/go-vscale"
)

func resourceRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceRecordCreate,
		Read:   resourceRecordRead,
		Exists: resourceRecordExists,
		Update: resourceRecordUpdate,
		Delete: resourceRecordDelete,

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
		},
	}
}

func resourceRecordCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*vscale.WebClient)

	domainID := d.Get("domain").(int)
	name := d.Get("name").(string)
	recordType := d.Get("type").(string)
	content := d.Get("content").(string)
	ttl := d.Get("ttl").(int)

	record, res, err := client.DomainRecord.Create(int64(domainID), name, recordType, int64(ttl), content)
	if err != nil {
		b, berr := ioutil.ReadAll(res.Body)
		if berr != nil {
			return errors.Wrap(err, "reading response body failed")
		}

		return errors.Wrapf(err, "creating domain record failed with body: %s", string(b))
	}

	d.SetId(strconv.FormatInt(record.ID, 10))

	return nil
}

func resourceRecordUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*vscale.WebClient)

	recordID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}
	domainID := d.Get("domain").(int)

	name := d.Get("name").(string)
	recordType := d.Get("type").(string)
	content := d.Get("content").(string)
	ttl := d.Get("ttl").(int)

	_, res, err := client.DomainRecord.Update(
		int64(domainID), recordID, name, recordType, int64(ttl), content,
	)
	if res.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil {
		return errors.Wrap(err, "updating domain record failed")
	}

	return nil
}

func resourceRecordRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*vscale.WebClient)

	recordID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	domainID := d.Get("domain").(int)

	record, res, err := client.DomainRecord.Get(int64(domainID), recordID)
	if res.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil {
		return errors.Wrap(err, "getting domain record failed")
	}

	d.Set("domain", domainID)
	d.Set("name", record.Name)
	d.Set("type", record.Type)
	d.Set("content", record.Content)
	d.Set("ttl", record.TTL)

	return nil
}

func resourceRecordExists(d *schema.ResourceData, m interface{}) (bool, error) {
	client := m.(*vscale.WebClient)

	recordID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return true, errors.Wrap(err, "invalid id")
	}

	domainID := d.Get("domain").(int)

	_, res, err := client.DomainRecord.Get(int64(domainID), recordID)
	if res.StatusCode == http.StatusNotFound {
		d.SetId("")
		return false, nil
	}
	if err != nil {
		return true, errors.Wrap(err, "getting domain record failed")
	}

	return true, nil
}

func resourceRecordDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*vscale.WebClient)

	recordID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	domainID := d.Get("domain").(int)

	_, res, err := client.DomainRecord.Remove(int64(domainID), recordID)
	if res.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil {
		return errors.Wrap(err, "removing domain record failed")
	}

	return nil
}
