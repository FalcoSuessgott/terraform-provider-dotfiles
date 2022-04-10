package provider

import (
	"fmt"

	"github.com/FalcoSuessgott/terraform-provider-dotfiles/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDotfilesLinkFile() *schema.Resource {
	return &schema.Resource{

		Create: resourceDotfilesLinkFileCreate,
		Read:   resourceDotfilesLinkFileRead,
		Update: resourceDotfilesLinkFileUpdate,
		Delete: resourceDotfilesLinkFileDelete,

		Schema: map[string]*schema.Schema{
			"src": {
				Type:        schema.TypeString,
				Description: "absolute source path of the file",
				Required:    true,
			},
			"dest": {
				Type:        schema.TypeString,
				Description: "absolute destination path of the file",
				Required:    true,
			},
			"content": {
				Type:        schema.TypeString,
				Description: "content of the file",
				Computed:    true,
			},
		},
	}
}

func resourceDotfilesLinkFileCreate(d *schema.ResourceData, meta interface{}) error {
	src := d.Get("src").(string)
	dest := d.Get("dest").(string)

	if !utils.Exists(src) {
		return fmt.Errorf("error file %s doesnt exist", src)
	}

	if err := utils.LinkFile(src, dest); err != nil {
		return fmt.Errorf("error linking %s to %s: %w", src, dest, err)
	}

	d.SetId(src)

	if err := d.Set("content", string(utils.ReadFile(src))); err != nil {
		return fmt.Errorf("error setting content: %w", err)
	}

	return nil
}

func resourceDotfilesLinkFileRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDotfilesLinkFileUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDotfilesLinkFileDelete(d *schema.ResourceData, meta interface{}) error {
	dest := d.Get("dest").(string)

	if err := utils.RemoveFile(dest); err != nil {
		return fmt.Errorf("error deleting %s: %w", dest, err)
	}

	return nil
}
