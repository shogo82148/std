// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package loadmacho implements a Mach-O file reader.
package loadmacho

import (
	"github.com/shogo82148/std/cmd/internal/bio"
	"github.com/shogo82148/std/cmd/internal/sys"
	"github.com/shogo82148/std/cmd/link/internal/loader"
)

// TODO(crawshaw): de-duplicate these symbols with cmd/link/internal/ld
const (
	MACHO_X86_64_RELOC_UNSIGNED = 0
	MACHO_X86_64_RELOC_SIGNED   = 1
	MACHO_ARM64_RELOC_ADDEND    = 10
)

// ldMachoSym.type_
const (
	N_EXT  = 0x01
	N_TYPE = 0x1e
	N_STAB = 0xe0
)

// ldMachoSym.desc
const (
	N_WEAK_REF = 0x40
	N_WEAK_DEF = 0x80
)

const (
	LdMachoCpuVax         = 1
	LdMachoCpu68000       = 6
	LdMachoCpu386         = 7
	LdMachoCpuAmd64       = 1<<24 | 7
	LdMachoCpuMips        = 8
	LdMachoCpu98000       = 10
	LdMachoCpuHppa        = 11
	LdMachoCpuArm         = 12
	LdMachoCpuArm64       = 1<<24 | 12
	LdMachoCpu88000       = 13
	LdMachoCpuSparc       = 14
	LdMachoCpu860         = 15
	LdMachoCpuAlpha       = 16
	LdMachoCpuPower       = 18
	LdMachoCmdSegment     = 1
	LdMachoCmdSymtab      = 2
	LdMachoCmdSymseg      = 3
	LdMachoCmdThread      = 4
	LdMachoCmdDysymtab    = 11
	LdMachoCmdSegment64   = 25
	LdMachoFileObject     = 1
	LdMachoFileExecutable = 2
	LdMachoFileFvmlib     = 3
	LdMachoFileCore       = 4
	LdMachoFilePreload    = 5
)

// Load the Mach-O file pn from f.
// Symbols are written into syms, and a slice of the text symbols is returned.
func Load(l *loader.Loader, arch *sys.Arch, localSymVersion int, f *bio.Reader, pkg string, length int64, pn string) (textp []loader.Sym, err error)
