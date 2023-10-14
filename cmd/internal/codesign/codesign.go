// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package codesign provides basic functionalities for
// ad-hoc code signing of Mach-O files.
//
// This is not a general tool for code-signing. It is made
// specifically for the Go toolchain. It uses the same
// ad-hoc signing algorithm as the Darwin linker.
package codesign

import (
	"github.com/shogo82148/std/debug/macho"
	"github.com/shogo82148/std/io"
)

const LC_CODE_SIGNATURE = 0x1d

const (
	CSMAGIC_REQUIREMENT        = 0xfade0c00
	CSMAGIC_REQUIREMENTS       = 0xfade0c01
	CSMAGIC_CODEDIRECTORY      = 0xfade0c02
	CSMAGIC_EMBEDDED_SIGNATURE = 0xfade0cc0
	CSMAGIC_DETACHED_SIGNATURE = 0xfade0cc1

	CSSLOT_CODEDIRECTORY = 0
)

const (
	CS_HASHTYPE_SHA1             = 1
	CS_HASHTYPE_SHA256           = 2
	CS_HASHTYPE_SHA256_TRUNCATED = 3
	CS_HASHTYPE_SHA384           = 4
)

const (
	CS_EXECSEG_MAIN_BINARY     = 0x1
	CS_EXECSEG_ALLOW_UNSIGNED  = 0x10
	CS_EXECSEG_DEBUGGER        = 0x20
	CS_EXECSEG_JIT             = 0x40
	CS_EXECSEG_SKIP_LV         = 0x80
	CS_EXECSEG_CAN_LOAD_CDHASH = 0x100
	CS_EXECSEG_CAN_EXEC_CDHASH = 0x200
)

type Blob struct {
	typ    uint32
	offset uint32
}

type SuperBlob struct {
	magic  uint32
	length uint32
	count  uint32
}

type CodeDirectory struct {
	magic         uint32
	length        uint32
	version       uint32
	flags         uint32
	hashOffset    uint32
	identOffset   uint32
	nSpecialSlots uint32
	nCodeSlots    uint32
	codeLimit     uint32
	hashSize      uint8
	hashType      uint8
	_pad1         uint8
	pageSize      uint8
	_pad2         uint32
	scatterOffset uint32
	teamOffset    uint32
	_pad3         uint32
	codeLimit64   uint64
	execSegBase   uint64
	execSegLimit  uint64
	execSegFlags  uint64
}

// CodeSigCmd is Mach-O LC_CODE_SIGNATURE load command.
type CodeSigCmd struct {
	Cmd      uint32
	Cmdsize  uint32
	Dataoff  uint32
	Datasize uint32
}

func FindCodeSigCmd(f *macho.File) (CodeSigCmd, bool)

// Size computes the size of the code signature.
// id is the identifier used for signing (a field in CodeDirectory blob, which
// has no significance in ad-hoc signing).
func Size(codeSize int64, id string) int64

// Sign generates an ad-hoc code signature and writes it to out.
// out must have length at least Size(codeSize, id).
// data is the file content without the signature, of size codeSize.
// textOff and textSize is the file offset and size of the text segment.
// isMain is true if this is a main executable.
// id is the identifier used for signing (a field in CodeDirectory blob, which
// has no significance in ad-hoc signing).
func Sign(out []byte, data io.Reader, id string, codeSize, textOff, textSize int64, isMain bool)
