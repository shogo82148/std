// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package loadxcoff implements a XCOFF file reader.
package loadxcoff

import (
	"github.com/shogo82148/std/cmd/internal/bio"
	"github.com/shogo82148/std/cmd/internal/sys"
	"github.com/shogo82148/std/cmd/link/internal/loader"
)

// loads the Xcoff file pn from f.
// Symbols are written into loader, and a slice of the text symbols is returned.
func Load(l *loader.Loader, arch *sys.Arch, localSymVersion int, input *bio.Reader, pkg string, length int64, pn string) (textp []loader.Sym, err error)
