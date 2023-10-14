// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modcmd

import (
	"github.com/shogo82148/std/context"

	"github.com/shogo82148/std/cmd/go/internal/modfetch/codehost"
)

// A ModuleJSON describes the result of go mod download.
type ModuleJSON struct {
	Path     string `json:",omitempty"`
	Version  string `json:",omitempty"`
	Query    string `json:",omitempty"`
	Error    string `json:",omitempty"`
	Info     string `json:",omitempty"`
	GoMod    string `json:",omitempty"`
	Zip      string `json:",omitempty"`
	Dir      string `json:",omitempty"`
	Sum      string `json:",omitempty"`
	GoModSum string `json:",omitempty"`

	Origin *codehost.Origin `json:",omitempty"`
	Reuse  bool             `json:",omitempty"`
}

// DownloadModule runs 'go mod download' for m.Path@m.Version,
// leaving the results (including any error) in m itself.
func DownloadModule(ctx context.Context, m *ModuleJSON)
