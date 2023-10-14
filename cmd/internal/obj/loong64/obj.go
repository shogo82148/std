// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loong64

import (
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/sys"
)

var Linkloong64 = obj.LinkArch{
	Arch:           sys.ArchLoong64,
	Init:           buildop,
	Preprocess:     preprocess,
	Assemble:       span0,
	Progedit:       progedit,
	DWARFRegisters: LOONG64DWARFRegisters,
}
