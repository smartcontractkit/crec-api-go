package generate

//go:generate go tool oapi-codegen -config ginserver/oapi-cfg.yaml ./spec/openapi.yaml
//go:generate go tool oapi-codegen -config stdserver/oapi-cfg.yaml ./spec/openapi.yaml
//go:generate go tool oapi-codegen -config client/oapi-cfg.yaml ./spec/openapi.yaml
