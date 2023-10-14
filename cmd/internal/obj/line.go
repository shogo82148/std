// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package obj

import (
	"github.com/shogo82148/std/cmd/internal/goobj"
)

// AddImport adds a package to the list of imported packages.
func (ctxt *Link) AddImport(pkg string, fingerprint goobj.FingerprintType)
