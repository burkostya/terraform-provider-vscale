package main

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	vscale "github.com/vganyn/go-vscale"
)

func resourceSSHKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSHKeyCreate,
		Read:   resourceSSHKeyRead,
		Delete: resourceSSHKeyDelete,
		Exists: resourceSSHKeyExists,

		Schema: map[string]*schema.Schema{
			// "id": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Computed: true,
			// },
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSSHKeyCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*vscale.WebClient)

	name := d.Get("name").(string)
	key := d.Get("key").(string)

	sshKey, _, err := client.SSHKey.Create(key, name)
	if err != nil {
		return errors.Wrap(err, "creating ssh key failed")
	}

	d.SetId(strconv.FormatInt(sshKey.ID, 10))

	return nil
}

func resourceSSHKeyRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*vscale.WebClient)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid SSH key id")
	}

	keys, _, err := client.SSHKey.List()
	if err != nil {
		return errors.Wrap(err, "listing ssh keys failed")
	}

	var sshKey vscale.SSHKey

	if keys != nil {
		for _, key := range *keys {
			if key.ID == id {
				sshKey = key
			}
		}
	}

	if sshKey.ID == 0 {
		d.SetId("")
		return nil
	}

	d.Set("key", sshKey.Key)
	d.Set("name", sshKey.Name)

	return nil
}

func resourceSSHKeyExists(d *schema.ResourceData, m interface{}) (bool, error) {
	client := m.(*vscale.WebClient)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return false, errors.Wrap(err, "invalid SSH key id")
	}

	keys, _, err := client.SSHKey.List()
	if err != nil {
		return true, errors.Wrap(err, "listing ssh keys failed")
	}

	if keys != nil {
		for _, key := range *keys {
			if key.ID == id {
				return true, nil
			}
		}
	}

	d.SetId("")
	return false, nil
}

func resourceSSHKeyDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*vscale.WebClient)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid SSH key id")
	}

	ok, res, err := client.SSHKey.Remove(id)
	if err != nil {
		return errors.Wrap(err, "removing SSH key failed")
	}
	if !ok && res.StatusCode != 200 {
		errText := fmt.Sprintf("removing SSH key failed, http code: %d", res.StatusCode)
		return errors.New(errText)
	}

	return nil
}
