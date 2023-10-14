// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go-specific code shared across loaders (5l, 6l, 8l).

package ld

import (
	"github.com/shogo82148/std/cmd/link/internal/loader"
)

func Adddynsym(ldr *loader.Loader, target *Target, syms *ArchSyms, s loader.Sym)
