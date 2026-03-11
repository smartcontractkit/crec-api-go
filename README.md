# CREC API Go

Go client, server stubs, and shared models for the CREC (CRE Connect) API. Generated from OpenAPI specifications.

## Overview

This repository provides:

- **client** — HTTP client for calling CREC API endpoints (channels, watchers, wallets, operations, events)
- **models** — Shared types such as VerifiableEvent, EVMEvent, Channel, Wallet, and operation/event types
- **ginserver** / **stdserver** — Generated server stubs for implementing the API (Gin or net/http)

## Installation

```bash
go get github.com/smartcontractkit/crec-api-go/client
go get github.com/smartcontractkit/crec-api-go/models
```

> Most consumers should use [crec-sdk](https://github.com/smartcontractkit/crec-sdk) instead of this package directly.

## Usage

### Client

```go
import (
    "context"
    "log"
    "github.com/smartcontractkit/crec-api-go/client"
)

func main() {
    c, err := client.NewClient("https://api.example.com")
    if err != nil {
        log.Fatal(err)
    }

    channels, err := c.GetChannels(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
    // Use channels...
}
```

Use `client.WithRequestEditorFn` to add auth headers (e.g. API key) to requests.

### Models

The models package is used by crec-workflow-utils, crec-sdk-ext-dta, crec-sdk-ext-dvp, and other CREC components for VerifiableEvent, EVMEvent, and related types.

## Regenerating Bindings

After modifying `api/openapi.yaml` or `models/models.yaml`:

```bash
make tools && make generate
```

This runs oapi-codegen for the client, server stubs, and models.

## Validating the OpenAPI Spec

```bash
make validate-openapi
```

## Related

- [crec-sdk](https://github.com/smartcontractkit/crec-sdk) — High-level SDK (most consumers should use this instead of the raw client)
- [crec-workflow-utils](https://github.com/smartcontractkit/crec-workflow-utils) — Shared utilities for event-listener workflows

## License

[MIT](LICENSE.md)
