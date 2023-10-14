// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO/NICETOHAVE:
//   - eliminate DW_CLS_ if not used
//   - package info in compilation units
//   - assign types to their packages
//   - gdb uses c syntax, meaning clumsy quoting is needed for go identifiers. eg
//     ptype struct '[]uint8' and qualifiers need to be quoted away
//   - file:line info for variables
//   - make strings a typedef so prettyprinters can see the underlying string type

package ld

import (
	"github.com/shogo82148/std/internal/abi"
)

// https://sourceware.org/gdb/onlinedocs/gdb/dotdebug_005fgdb_005fscripts-section.html
// Each entry inside .debug_gdb_scripts section begins with a non-null prefix
// byte that specifies the kind of entry. The following entries are supported:
const (
	GdbScriptPythonFileId = 1
	GdbScriptSchemeFileId = 3
	GdbScriptPythonTextId = 4
	GdbScriptSchemeTextId = 6
)

// synthesizemaptypes is way too closely married to runtime/hashmap.c
const (
	MaxKeySize = abi.MapMaxKeyBytes
	MaxValSize = abi.MapMaxElemBytes
	BucketSize = abi.MapBucketCount
)

/*
 * Generate a sequence of opcodes that is as short as possible.
 * See section 6.2.5
 */
const (
	LINE_BASE   = -4
	LINE_RANGE  = 10
	PC_RANGE    = (255 - OPCODE_BASE) / LINE_RANGE
	OPCODE_BASE = 11
)

const (
	COMPUNITHEADERSIZE = 4 + 2 + 4 + 1
)
