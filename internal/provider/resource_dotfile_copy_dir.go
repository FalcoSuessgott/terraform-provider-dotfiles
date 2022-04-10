package provider

import (
	"fmt"

	"github.com/FalcoSuessgott/terraform-provider-dotfiles/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDotfilesCopyDir() *schema.Resource {
	return &schema.Resource{

		Create: resourceDotfilesCopyDirCreate,
		Read:   resourceDotfilesCopyDirRead,
		Update: resourceDotfilesCopyDirUpdate,
		Delete: resourceDotfilesCopyDirDelete,

		Schema: map[string]*schema.Schema{
			"src": {
				Type:        schema.TypeString,
				Description: "absolute source path of the directory",
				Required:    true,
			},
			"dest": {
				Type:        schema.TypeString,
				Description: "absolute destination path of the directory",
				Required:    true,
			},
			"permissions": {
				Type:        schema.TypeString,
				Description: "permissions of the destination file or directory (defaults to 755)",
				Optional:    true,
				Default:     "755",
			},
		},
	}
}

func resourceDotfilesCopyDirCreate(d *schema.ResourceData, meta interface{}) error {
	// input
	src := d.Get("src").(string)
	dest := d.Get("dest").(string)
	// permissions := d.Get("permissions").(string)

	// verify
	if !utils.Exists(src) {
		return fmt.Errorf("error directory %s doesnt exist", src)
	}

	if err := utils.CopyDir(src, dest); err != nil {
		return fmt.Errorf("error copying directory %s to %s: %w", src, dest, err)
	}

	d.SetId(src)

	return nil
}

func resourceDotfilesCopyDirRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDotfilesCopyDirUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDotfilesCopyDirDelete(d *schema.ResourceData, meta interface{}) error {
	dest := d.Get("dest").(string)

	if err := utils.RemoveFile(dest); err != nil {
		return fmt.Errorf("error deleting %s: %w", dest, err)
	}

	return nil
}
