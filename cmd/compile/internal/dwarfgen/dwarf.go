// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dwarfgen

import (
	"github.com/shogo82148/std/cmd/internal/dwarf"
	"github.com/shogo82148/std/cmd/internal/obj"
)

func Info(fnsym *obj.LSym, infosym *obj.LSym, curfn obj.Func) (scopes []dwarf.Scope, inlcalls dwarf.InlCalls)

// RecordFlags records the specified command-line flags to be placed
// in the DWARF info.
func RecordFlags(flags ...string)

// RecordPackageName records the name of the package being
// compiled, so that the linker can save it in the compile unit's DIE.
func RecordPackageName()
