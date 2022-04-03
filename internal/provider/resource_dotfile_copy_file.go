package provider

import (
	"fmt"

	"github.com/FalcoSuessgott/terraform-provider-dotfiles/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDotfilesCopyFile() *schema.Resource {
	return &schema.Resource{

		Create: resourceDotfilesCopyFileCreate,
		Read:   resourceDotfilesCopyFileRead,
		Update: resourceDotfilesCopyFileUpdate,
		Delete: resourceDotfilesCopyFileDelete,

		Schema: map[string]*schema.Schema{
			"src": {
				Type:         schema.TypeString,
				Description:  "absolute source path of the file",
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"src", "content"},
			},
			"content": {
				Type:         schema.TypeString,
				Description:  "content of the file.",
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"src", "content"},
			},
			"dest": {
				Type:        schema.TypeString,
				Description: "absolute destination path of the file",
				Required:    true,
			},
			"permissions": {
				Type:        schema.TypeInt,
				Description: "permissions of the destination file or file (defaults to 755)",
				Optional:    true,
				Default:     755,
			},
		},
	}
}

func resourceDotfilesCopyFileCreate(d *schema.ResourceData, meta interface{}) error {
	src := d.Get("src").(string)
	dest := d.Get("dest").(string)
	content := d.Get("content").(string)
	permissions := d.Get("permissions").(int)

	switch {
	case src != "":
		if !utils.Exists(src) {
			return fmt.Errorf("error file %s doesnt exist", src)
		}

		if err := utils.CopyFile(src, dest, permissions); err != nil {
			return fmt.Errorf("error copying file %s to %s: %w", src, dest, err)
		}
	case content != "":
		if err := utils.WriteFile([]byte(content), dest, permissions); err != nil {
			return fmt.Errorf("error copying file %s to %s: %w", src, dest, err)
		}
	default:
		return fmt.Errorf("this shouldnt happen")
	}

	d.SetId(utils.MD5Sum([]byte(content)))

	return resourceDotfilesCopyFileRead(d, meta)
}

func resourceDotfilesCopyFileRead(d *schema.ResourceData, meta interface{}) error {
	if err := d.Set("content", d.Get("content").(string)); err != nil {
		return fmt.Errorf("error setting content: %w", err)
	}

	return nil
}

func resourceDotfilesCopyFileUpdate(d *schema.ResourceData, meta interface{}) error {
	src := d.Get("src").(string)
	dest := d.Get("dest").(string)
	content := d.Get("content").(string)
	permissions := d.Get("permissions").(int)

	if d.HasChange("src") {
		if !utils.Exists(src) {
			return fmt.Errorf("error file %s doesnt exist", src)
		}

		if err := utils.CopyFile(src, dest, permissions); err != nil {
			return fmt.Errorf("error copying file %s to %s: %w", src, dest, err)
		}

		content = string(utils.ReadFile(src))
	}

	if d.HasChange("content") {
		if err := utils.WriteFile([]byte(content), dest, permissions); err != nil {
			return fmt.Errorf("error copying file %s to %s: %w", src, dest, err)
		}
	}

	return resourceDotfilesCopyFileRead(d, meta)
}

func resourceDotfilesCopyFileDelete(d *schema.ResourceData, meta interface{}) error {
	dest := d.Get("dest").(string)

	if err := utils.RemoveFile(dest); err != nil {
		return fmt.Errorf("error deleting %s: %w", dest, err)
	}

	return nil
}
