// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Gopls (pronounced “go please”) is an LSP server for Go.
// The Language Server Protocol allows any text editor
// to be extended with IDE-like features;
// see https://langserver.org/ for details.
//
// See https://github.com/golang/tools/blob/master/gopls/README.md
// for the most up-to-date documentation.
package main // import "golang.org/x/tools/gopls"

//go:generate go run doc/generate.go

import (
	"context"
	"os"

	"golang.org/x/telemetry/counter"
	"golang.org/x/tools/gopls/internal/cmd"
	"golang.org/x/tools/gopls/internal/hooks"
	versionpkg "golang.org/x/tools/gopls/internal/version"
	"golang.org/x/tools/internal/tool"
)

var version = "" // if set by the linker, overrides the gopls version

func main() {
	versionpkg.VersionOverride = version

	counter.Open() // Enable telemetry counter writing.
	ctx := context.Background()
	tool.Main(ctx, cmd.New(hooks.Options), os.Args[1:])
}
