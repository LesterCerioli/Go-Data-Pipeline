// Package docs embeds the OpenAPI descriptions served by the API.
package docs

import _ "embed"

// SwaggerJSON and SwaggerYAML contain the API's Swagger 2.0 specification.
//
//go:embed swagger.json
var SwaggerJSON []byte

//go:embed swagger.yaml
var SwaggerYAML []byte
