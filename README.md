# terraform-provider-dotfiles

## *still work in progress*

An experimental terraform provider for linking and copying files across the filesystem, aimed to be used as a dotfile manager


# Why
* terraform comes with a clean and flexible HCL [syntax](https://www.terraform.io/language/syntax/configuration) including lots in integrations like vscode-plugins, linter etc.
* rich featureset of predefined [functions](https://www.terraform.io/language/functions) useful for customizing and automating your dotfile setup
* can be combined with more than 1000 other useful [provider](https://registry.terraform.io/), such as [docker](https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs), [restapi](https://registry.terraform.io/providers/Mastercard/restapi/1.16.2) and many other
* [state](https://www.terraform.io/language/state) based, which can be used to deploy your dotfiles across many machines and reset them to a certain state

# Example

```hcl
terraform {
  required_providers {
    dotfiles = {
      source  = "FalcoSuessgott/dotfiles"
      version = "0.0.1"
    }
  }
}

provider "dotfiles" {}

resource "dotfiles_link_file" "bashrc" {
  src  = abspath("configs/bashrc")
  dest = pathexpand("~/.bashrc")
}

resource "dotfiles_copy_file" "zshrc" {
  src  = abspath("configs/zshrc")
  dest = pathexpand("~/.zshrc")
}

resource "dotfiles_copy_file" "aliases" {
  content = <<EOT
alias ll="ls -la"
EOT

  dest = pathexpand("~/.bash_aliases")
}
```