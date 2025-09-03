# CVN-API :: Golang Bindings

## Prerequisites

Before running code generation, make sure you have the following tools installed:

```bash
go install github.com/atombender/go-jsonschema@latest
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

To regenerate the Golang bindings for the CVN-API, run:

```bash
make generate
```