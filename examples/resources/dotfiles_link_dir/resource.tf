resource "dotfiles_link_dir" "zshrc" {
  src  = abspath("configs/.zshrc")
  dest = abspath(".zshrc")
  permissions = "644"
}