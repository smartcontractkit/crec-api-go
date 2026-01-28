package generate

//go:generate go tool oapi-codegen -config ginserver/oapi-cfg.yaml ./api/openapi.yaml
//go:generate go tool oapi-codegen -config stdserver/oapi-cfg.yaml ./api/openapi.yaml
//go:generate go tool oapi-codegen -config client/oapi-cfg.yaml ./api/openapi.yaml
//go:generate go tool oapi-codegen -config models/oapi-cfg.yaml ./models/models.yaml
