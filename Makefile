.PHONY: tools
tools:
	go install github.com/atombender/go-jsonschema@latest
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	go install github.com/daveshanley/vacuum@latest

.PHONY: generate
generate:
	go generate ./
	go-jsonschema services/dvp/schema/dvp.json -p events -o services/dvp/gen/events/events.gen.go -t
	abigen --abi services/dvp/abi/CCIPDVPCoordinator.abi.json --pkg contract --out services/dvp/gen/contract/contract.gen.go
	abigen --abi services/ccip/abi/IRouterClient.abi.json --pkg routerclient --out services/ccip/gen/routerclient/routerclient.gen.go
	abigen --abi services/accounts/abi/AccountFactory.abi.json --pkg accounts --out services/accounts/gen/accounts/accounts.gen.go

.PHONY: validate-openapi
validate-openapi:
	vacuum lint --ruleset .vacuum.yaml ./spec/openapi.yaml

