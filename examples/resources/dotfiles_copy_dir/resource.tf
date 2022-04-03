resource "dotfiles_copy_dir" "aliases" {
  src  = abspath("configs/aliases")
  dest = abspath("aliases")
  permissions = "644"
}