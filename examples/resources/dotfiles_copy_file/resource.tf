resource "dotfiles_copy_file" "zshrc" {
  src  = abspath("configs/.zshrc")
  dest = abspath(".zshrc")
  permissions = "644"
}