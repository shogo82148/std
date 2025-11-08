// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"github.com/shogo82148/std/cmd/internal/sys"
	"github.com/shogo82148/std/cmd/link/internal/loader"
	"github.com/shogo82148/std/debug/elf"
)

// ElfEhdr is the ELF file header.
type ElfEhdr elf.Header64

// ElfShdr is an ELF section entry, plus the section index.
type ElfShdr struct {
	elf.Section64
	shnum elf.SectionIndex
}

// ElfPhdr is the ELF program, or segment, header.
type ElfPhdr elf.ProgHeader

const (
	ELF64HDRSIZE  = 64
	ELF64PHDRSIZE = 56
	ELF64SHDRSIZE = 64
	ELF64RELSIZE  = 16
	ELF64RELASIZE = 24
	ELF64SYMSIZE  = 24
	ELF32HDRSIZE  = 52
	ELF32PHDRSIZE = 32
	ELF32SHDRSIZE = 40
	ELF32SYMSIZE  = 16
	ELF32RELSIZE  = 8
)

// ELFRESERVE is the total amount of space to reserve at the
// start of the file for Header, PHeaders, SHeaders, and interp.
// May waste some space.
// On FreeBSD, cannot be larger than a page.
const ELFRESERVE = 4096

const (
	NSECT = 400
)

var (
	Nelfsym = 1
)

// ELFArch includes target-specific hooks for ELF targets.
// This is initialized by the target-specific Init function
// called by the linker's main function in cmd/link/main.go.
type ELFArch struct {
	Androiddynld   string
	Linuxdynld     string
	LinuxdynldMusl string
	Freebsddynld   string
	Netbsddynld    string
	Openbsddynld   string
	Dragonflydynld string
	Solarisdynld   string

	Reloc1    func(*Link, *OutBuf, *loader.Loader, loader.Sym, loader.ExtReloc, int, int64) bool
	RelocSize uint32
	SetupPLT  func(ctxt *Link, ldr *loader.Loader, plt, gotplt *loader.SymbolBuilder, dynamic loader.Sym)

	// DynamicReadOnly can be set to true to make the .dynamic
	// section read-only. By default it is writable.
	// This is used by MIPS targets.
	DynamicReadOnly bool
}

type Elfstring struct {
	s   string
	off int
}

// Elfinit initializes the global ehdr variable that holds the ELF header.
// It will be updated as write section and program headers.
func Elfinit(ctxt *Link)

func Elfwritedynent(arch *sys.Arch, s *loader.SymbolBuilder, tag elf.DynTag, val uint64)

func Elfwritedynentsymplus(ctxt *Link, s *loader.SymbolBuilder, tag elf.DynTag, t loader.Sym, add int64)

// member of .gnu.attributes of MIPS for fpAbi
const (
	// No floating point is present in the module (default)
	MIPS_FPABI_NONE = 0
	// FP code in the module uses the FP32 ABI for a 32-bit ABI
	MIPS_FPABI_ANY = 1
	// FP code in the module only uses single precision ABI
	MIPS_FPABI_SINGLE = 2
	// FP code in the module uses soft-float ABI
	MIPS_FPABI_SOFT = 3
	// FP code in the module assumes an FPU with FR=1 and has 12
	// callee-saved doubles. Historic, no longer supported.
	MIPS_FPABI_HIST = 4
	// FP code in the module uses the FPXX  ABI
	MIPS_FPABI_FPXX = 5
	// FP code in the module uses the FP64  ABI
	MIPS_FPABI_FP64 = 6
	// FP code in the module uses the FP64A ABI
	MIPS_FPABI_FP64A = 7
)

// NetBSD Signature (as per sys/exec_elf.h)
const (
	ELF_NOTE_NETBSD_NAMESZ  = 7
	ELF_NOTE_NETBSD_DESCSZ  = 4
	ELF_NOTE_NETBSD_TAG     = 1
	ELF_NOTE_NETBSD_VERSION = 700000000
)

var ELF_NOTE_NETBSD_NAME = []byte("NetBSD\x00")

// OpenBSD Signature
const (
	ELF_NOTE_OPENBSD_NAMESZ  = 8
	ELF_NOTE_OPENBSD_DESCSZ  = 4
	ELF_NOTE_OPENBSD_TAG     = 1
	ELF_NOTE_OPENBSD_VERSION = 0
)

var ELF_NOTE_OPENBSD_NAME = []byte("OpenBSD\x00")

// FreeBSD Signature (as per sys/elf_common.h)
const (
	ELF_NOTE_FREEBSD_NAMESZ            = 8
	ELF_NOTE_FREEBSD_DESCSZ            = 4
	ELF_NOTE_FREEBSD_ABI_TAG           = 1
	ELF_NOTE_FREEBSD_NOINIT_TAG        = 2
	ELF_NOTE_FREEBSD_FEATURE_CTL_TAG   = 4
	ELF_NOTE_FREEBSD_VERSION           = 1203000
	ELF_NOTE_FREEBSD_FCTL_ASLR_DISABLE = 0x1
)

const ELF_NOTE_FREEBSD_NAME = "FreeBSD\x00"

// Build info note
const (
	ELF_NOTE_BUILDINFO_NAMESZ = 4
	ELF_NOTE_BUILDINFO_TAG    = 3
)

var ELF_NOTE_BUILDINFO_NAME = []byte("GNU\x00")

// Go specific notes
const (
	ELF_NOTE_GOPKGLIST_TAG = 1
	ELF_NOTE_GOABIHASH_TAG = 2
	ELF_NOTE_GODEPS_TAG    = 3
	ELF_NOTE_GOBUILDID_TAG = 4
)

var ELF_NOTE_GO_NAME = []byte("Go\x00\x00")

type Elfaux struct {
	next *Elfaux
	num  int
	vers string
}

type Elflib struct {
	next *Elflib
	aux  *Elfaux
	file string
}

func Asmbelfsetup()
