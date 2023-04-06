//go:build tools
// +build tools

// Tools is a dummy package to force go mod to download the tools we need to build the project.
package tools

import (
	_ "github.com/mgechev/revive"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
