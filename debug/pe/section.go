// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pe

import (
	"github.com/shogo82148/std/io"
)

// SectionHeader32は、実際のPE COFFセクションヘッダーを表します。
type SectionHeader32 struct {
	Name                 [8]uint8
	VirtualSize          uint32
	VirtualAddress       uint32
	SizeOfRawData        uint32
	PointerToRawData     uint32
	PointerToRelocations uint32
	PointerToLineNumbers uint32
	NumberOfRelocations  uint16
	NumberOfLineNumbers  uint16
	Characteristics      uint32
}

// Relocは、PE COFFの再配置を表します。
// 各セクションには独自の再配置リストが含まれています。
type Reloc struct {
	VirtualAddress   uint32
	SymbolTableIndex uint32
	Type             uint16
}

// SectionHeaderは、NameフィールドがGoの文字列に置き換えられたSectionHeader32と似ています。
type SectionHeader struct {
	Name                 string
	VirtualSize          uint32
	VirtualAddress       uint32
	Size                 uint32
	Offset               uint32
	PointerToRelocations uint32
	PointerToLineNumbers uint32
	NumberOfRelocations  uint16
	NumberOfLineNumbers  uint16
	Characteristics      uint32
}

// Sectionは、PE COFFセクションへのアクセスを提供します。
type Section struct {
	SectionHeader
	Relocs []Reloc

	// Embed ReaderAt for ReadAt method.
	// Do not embed SectionReader directly
	// to avoid having Read and Seek.
	// If a client wants Read and Seek it must use
	// Open() to avoid fighting over the seek offset
	// with other clients.
	io.ReaderAt
	sr *io.SectionReader
}

// Dataは、PEセクションsの内容を読み取り、返します。
//
// s.Offsetが0の場合、セクションには内容がなく、
// Dataは常に非nilのエラーを返します。
func (s *Section) Data() ([]byte, error)

// Openは、PEセクションsを読み取る新しいReadSeekerを返します。
//
// s.Offsetが0の場合、セクションには内容がなく、
// 返されたリーダーへのすべての呼び出しは非nilのエラーを返します。
func (s *Section) Open() io.ReadSeeker

// セクション特性フラグ。
const (
	IMAGE_SCN_CNT_CODE               = 0x00000020
	IMAGE_SCN_CNT_INITIALIZED_DATA   = 0x00000040
	IMAGE_SCN_CNT_UNINITIALIZED_DATA = 0x00000080
	IMAGE_SCN_LNK_COMDAT             = 0x00001000
	IMAGE_SCN_MEM_DISCARDABLE        = 0x02000000
	IMAGE_SCN_MEM_EXECUTE            = 0x20000000
	IMAGE_SCN_MEM_READ               = 0x40000000
	IMAGE_SCN_MEM_WRITE              = 0x80000000
)
