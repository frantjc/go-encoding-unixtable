GO = go
GOLANGCI-LINT = golangci-lint

fmt generate test:
	@$(GO) $@ ./...

download vendor verify:
	@$(GO) mod $@

lint:
	@$(GOLANGCI-LINT) run --fix

gen: generate
dl: download
ven: vendor
ver: verify
format: fmt

.PHONY: fmt generate test download vendor verify lint gen dl ven ver format
