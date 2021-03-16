NAME=hashicups
BINARY=packer-plugin-${NAME}

.PHONY: install-plugin run-example

build:
	@go build -o ${BINARY}

install: build
	@mkdir -p ~/.packer.d/plugins/
	@mv ${BINARY} ~/.packer.d/plugins/${BINARY}

run-example: install
	@packer build ./example

run-product-api:
	@cd example/product_api && docker-compose up -d

testacc: install
	@PACKER_ACC=1 go test -count 1 -v ./... -timeout=120m