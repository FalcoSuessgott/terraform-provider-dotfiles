package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// New returns the provider.
func New() func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{},
			ResourcesMap: map[string]*schema.Resource{
				"dotfiles_link_file": resourceDotfilesLinkFile(),
				"dotfiles_link_dir":  resourceDotfilesLinkDir(),
				"dotfiles_copy_file": resourceDotfilesCopyFile(),
				"dotfiles_copy_dir":  resourceDotfilesCopyDir(),
			},
			Schema: map[string]*schema.Schema{},
		}

		p.ConfigureContextFunc = nil

		return p
	}
}
