package main

//go:generate go tool oapi-codegen -config ginserver/oapi-cfg.yaml openapi.yaml
//go:generate go tool oapi-codegen -config stdserver/oapi-cfg.yaml openapi.yaml
//go:generate go tool oapi-codegen -config client/oapi-cfg.yaml openapi.yaml
