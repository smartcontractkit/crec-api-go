package spec

import (
	"embed"
	"io/fs"
)

// Embed the openapi.yaml file located at the package root
//
//go:embed openapi.yaml
var openapiFS embed.FS

func GetOpenAPISpec() (fs.File, error) {
	return openapiFS.Open("openapi.yaml")
}
