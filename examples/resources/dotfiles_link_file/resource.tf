resource "dotfiles_link_file" "zshrc" {
  src  = abspath("configs/.zshrc")
  dest = abspath(".zshrc")
  permissions = "644"
}