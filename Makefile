.PHONY: tools
tools:
	go install github.com/daveshanley/vacuum@latest

.PHONY: generate
generate:
	go generate ./

.PHONY: validate-openapi
validate-openapi:
	vacuum lint --ruleset .vacuum.yaml ./api/openapi.yaml
