TF_PLUGIN_DIR=~/.terraform.d/plugins
TF_PROVIDER_NAMESPACE=localhost/dev
TF_PROVIDER_PREFIX=terraform-provider
TF_PROVIDER_NAME=dotfiles
TF_PROVIDER_FULLNAME=$(TF_PROVIDER_PREFIX)-$(TF_PROVIDER_NAME)
TF_PROVIDER_VERSION=0.0.1
TF_PROVIDER_DIR=$(TF_PLUGIN_DIR)/$(TF_PROVIDER_NAMESPACE)/$(TF_PROVIDER_NAME)/$(TF_PROVIDER_VERSION)/linux_amd64

default: help

.PHONY: help
help: ## list makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: ## install local provider
	mkdir -p $(TF_PROVIDER_DIR)
	go build -o $(TF_PROVIDER_FULLNAME) .
	ln -sfv $(PWD)/$(TF_PROVIDER_FULLNAME) $(TF_PROVIDER_DIR)/$(TF_PROVIDER_FULLNAME)

.PHONY: fmt
fmt: ## format code
	gofumpt -l -w .
	gci -w .
	terraform -chdir="terraform/" fmt -recursive -write

.PHONY: tf
tf: clean install ## run tf suite
	terraform -chdir="terraform/" init -reconfigure
	terraform -chdir="terraform/" apply -auto-approve

PHONY: lint
lint: ## lint go files
	golangci-lint run -c .golang-ci.yml

.PHONY: docs
docs: ## generate provider docs
	tfplugindocs

.PHONY: destroy
destroy: ## run terraform destroy
	terraform -chdir="terraform/" destroy -auto-approve

.PHONY: clean
clean: ## clean tmp files
	rm -rf terraform/.terraform terraform/terraform.tfstate.backup terraform/.terraform.lock.hcl terraform/terraform.tfstate $(TF_PROVIDER_FULLNAME)
