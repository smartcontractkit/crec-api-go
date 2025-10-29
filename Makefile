.PHONY: tools
tools:
	go install github.com/atombender/go-jsonschema@latest
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest

.PHONY: generate
generate:
	go generate ./
	go-jsonschema services/dvp/schema/dvp.json -p events -o services/dvp/gen/events/events.gen.go -t
	go-jsonschema services/dta/schema/v1.0/dta.json -p events -o services/dta/gen/events/events.gen.go -t
	abigen --abi services/dvp/abi/CCIPDVPCoordinator.abi.json --pkg contract --out services/dvp/gen/contract/contract.gen.go
	abigen --abi services/ccip/abi/IRouterClient.abi.json --pkg routerclient --out services/ccip/gen/routerclient/routerclient.gen.go
	abigen --abi services/dta/abi/DTAOpenMarketplaceU.abi.json --pkg dtaopenmarketplace --out services/dta/gen/dtaopenmarketplace/dtaopenmarketplace.gen.go
	abigen --abi services/dta/abi/DTAWalletU.abi.json --pkg dtawallet --out services/dta/gen/dtawallet/dtawallet.gen.go
	abigen --abi services/dta/abi/DTARequestManagementU.abi.json --pkg dtarequestmanagement --out services/dta/gen/dtarequestmanagement/dtarequestmanagement.gen.go
	abigen --abi services/dta/abi/DTARequestSettlement.abi.json --pkg dtarequestsettlement --out services/dta/gen/dtarequestsettlement/dtarequestsettlement.gen.go
	abigen --abi services/wallets/abi/WalletFactory.abi.json --pkg wallets --out services/wallets/gen/wallets/wallets.gen.go

