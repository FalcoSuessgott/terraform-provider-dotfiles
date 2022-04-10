terraform {
  required_providers {
    dotfiles = {
      source  = "localhost/dev/dotfiles"
      version = "0.0.1"
    }
  }
}

provider "dotfiles" {}

locals {
  src_file = abspath("testfile.txt")
  src_dir  = abspath("testdir")
  dest_dir = "./result"
}

resource "dotfiles_link_file" "this" {
  src  = local.src_file
  dest = abspath("testfile_link.txt")
}

resource "dotfiles_copy_file" "this" {
  src         = local.src_file
  dest        = abspath("testfile_copy.txt")
  permissions = 0600
}

resource "dotfiles_copy_file" "content" {
  content = <<EOT
This is a test
EOT

  dest        = abspath("testfile_copy_content.txt")
  permissions = 0600
}

resource "dotfiles_link_dir" "this" {
  src  = local.src_dir
  dest = abspath("link")
}

resource "dotfiles_copy_dir" "this" {
  src  = local.src_dir
  dest = abspath("copy")
}