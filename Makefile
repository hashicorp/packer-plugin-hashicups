.PHONY: install-plugin run-example

install-plugin:
	@go build . && cp packer-plugin-hashicups ~/.packer.d/plugins/packer-plugin-hashicups

run-example: install-plugin
	@packer build ./example

acc-test: install-plugin
	@PACKER_ACC=1 go test -count 1 -v ./... -timeout=120m