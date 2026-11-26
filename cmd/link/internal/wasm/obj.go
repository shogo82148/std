// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wasm

import (
	"github.com/shogo82148/std/cmd/internal/sys"
	"github.com/shogo82148/std/cmd/link/internal/ld"
)

// Init returns the architecture configuration for the wasm linker.
func Init(arch *sys.Arch) (*sys.Arch, ld.Arch)
