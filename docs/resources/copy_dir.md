---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dotfiles_copy_dir Resource - terraform-provider-dotfiles"
subcategory: ""
description: |-
  
---

# dotfiles_copy_dir (Resource)



## Example Usage

```terraform
resource "dotfiles_copy_dir" "aliases" {
  src  = abspath("configs/aliases")
  dest = abspath("aliases")
  permissions = "644"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **dest** (String) absolute destination path of the directory
- **src** (String) absolute source path of the directory

### Optional

- **id** (String) The ID of this resource.
- **permissions** (String) permissions of the destination file or directory (defaults to 755)


