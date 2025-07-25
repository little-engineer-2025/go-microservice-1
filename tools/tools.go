//go:build tools
// +build tools

package main

import (
	_ "github.com/achiku/planter"
	_ "github.com/atombender/go-jsonschema/cmd/gojsonschema"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/loov/goda"
	_ "github.com/mikefarah/yq/v4"
	_ "github.com/vektra/mockery/v2"
)
