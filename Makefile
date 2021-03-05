.PHONY: install-plugin run-example

install-plugin:
	@go build . && cp packer-plugin-hashicups ~/.packer.d/plugins/packer-plugin-hashicups

run-example:
	@packer build ./example