VACUUM_VERSION ?= 0.20.5

.PHONY: tools
tools:
	go install github.com/daveshanley/vacuum@v$(VACUUM_VERSION)

.PHONY: generate
generate:
	go generate ./

.PHONY: validate-openapi
validate-openapi:
	vacuum lint --ruleset .vacuum.yaml ./api/openapi.yaml
