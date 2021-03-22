NAME=hashicups
BINARY=packer-plugin-${NAME}

.PHONY: install run-example

build:
	@go build -o ${BINARY}

install: build
	@mkdir -p ~/.packer.d/plugins/
	@mv ${BINARY} ~/.packer.d/plugins/${BINARY}

run-example: install
	@packer build ./example

run-hashicups-api:
	@cd example/hashicups_api && docker-compose up -d

testacc: install
	@PACKER_ACC=1 go test -count 1 -v ./... -timeout=120m