// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Mach-O header data structures
// Originally at:
// http://developer.apple.com/mac/library/documentation/DeveloperTools/Conceptual/MachORuntime/Reference/reference.html (since deleted by Apple)
// Archived copy at:
// https://web.archive.org/web/20090819232456/http://developer.apple.com/documentation/DeveloperTools/Conceptual/MachORuntime/index.html
// For cloned PDF see:
// https://github.com/aidansteele/osx-abi-macho-file-format-reference

package macho

// FileHeaderは、Mach-Oファイルヘッダーを表します。
type FileHeader struct {
	Magic  uint32
	Cpu    Cpu
	SubCpu uint32
	Type   Type
	Ncmd   uint32
	Cmdsz  uint32
	Flags  uint32
}

const (
	Magic32  uint32 = 0xfeedface
	Magic64  uint32 = 0xfeedfacf
	MagicFat uint32 = 0xcafebabe
)

// Typeは、Mach-Oファイルのタイプを表します。例えば、オブジェクトファイル、実行可能ファイル、または動的ライブラリなどです。
type Type uint32

const (
	TypeObj    Type = 1
	TypeExec   Type = 2
	TypeDylib  Type = 6
	TypeBundle Type = 8
)

func (t Type) String() string
func (t Type) GoString() string

// Cpuは、Mach-OのCPUタイプを表します。
type Cpu uint32

const (
	Cpu386   Cpu = 7
	CpuAmd64 Cpu = Cpu386 | cpuArch64
	CpuArm   Cpu = 12
	CpuArm64 Cpu = CpuArm | cpuArch64
	CpuPpc   Cpu = 18
	CpuPpc64 Cpu = CpuPpc | cpuArch64
)

func (i Cpu) String() string
func (i Cpu) GoString() string

// LoadCmdは、Mach-Oのロードコマンドです。
type LoadCmd uint32

const (
	LoadCmdSegment    LoadCmd = 0x1
	LoadCmdSymtab     LoadCmd = 0x2
	LoadCmdThread     LoadCmd = 0x4
	LoadCmdUnixThread LoadCmd = 0x5
	LoadCmdDysymtab   LoadCmd = 0xb
	LoadCmdDylib      LoadCmd = 0xc
	LoadCmdDylinker   LoadCmd = 0xf
	LoadCmdSegment64  LoadCmd = 0x19
	LoadCmdRpath      LoadCmd = 0x8000001c
)

func (i LoadCmd) String() string
func (i LoadCmd) GoString() string

type (
	// Segment32は、32ビットMach-Oのセグメントロードコマンドです。
	Segment32 struct {
		Cmd     LoadCmd
		Len     uint32
		Name    [16]byte
		Addr    uint32
		Memsz   uint32
		Offset  uint32
		Filesz  uint32
		Maxprot uint32
		Prot    uint32
		Nsect   uint32
		Flag    uint32
	}

	// Segment64は、64ビットMach-Oのセグメントロードコマンドです。
	Segment64 struct {
		Cmd     LoadCmd
		Len     uint32
		Name    [16]byte
		Addr    uint64
		Memsz   uint64
		Offset  uint64
		Filesz  uint64
		Maxprot uint32
		Prot    uint32
		Nsect   uint32
		Flag    uint32
	}

	// SymtabCmdは、Mach-Oのシンボルテーブルコマンドです。
	SymtabCmd struct {
		Cmd     LoadCmd
		Len     uint32
		Symoff  uint32
		Nsyms   uint32
		Stroff  uint32
		Strsize uint32
	}

	// DysymtabCmdは、Mach-Oの動的シンボルテーブルコマンドです。
	DysymtabCmd struct {
		Cmd            LoadCmd
		Len            uint32
		Ilocalsym      uint32
		Nlocalsym      uint32
		Iextdefsym     uint32
		Nextdefsym     uint32
		Iundefsym      uint32
		Nundefsym      uint32
		Tocoffset      uint32
		Ntoc           uint32
		Modtaboff      uint32
		Nmodtab        uint32
		Extrefsymoff   uint32
		Nextrefsyms    uint32
		Indirectsymoff uint32
		Nindirectsyms  uint32
		Extreloff      uint32
		Nextrel        uint32
		Locreloff      uint32
		Nlocrel        uint32
	}

	// DylibCmdは、Mach-Oの動的ライブラリロードコマンドです。
	DylibCmd struct {
		Cmd            LoadCmd
		Len            uint32
		Name           uint32
		Time           uint32
		CurrentVersion uint32
		CompatVersion  uint32
	}

	// RpathCmdは、Mach-Oのrpathコマンドです。
	RpathCmd struct {
		Cmd  LoadCmd
		Len  uint32
		Path uint32
	}

	// Threadは、Mach-Oのスレッド状態コマンドです。
	Thread struct {
		Cmd  LoadCmd
		Len  uint32
		Type uint32
		Data []uint32
	}
)

const (
	FlagNoUndefs              uint32 = 0x1
	FlagIncrLink              uint32 = 0x2
	FlagDyldLink              uint32 = 0x4
	FlagBindAtLoad            uint32 = 0x8
	FlagPrebound              uint32 = 0x10
	FlagSplitSegs             uint32 = 0x20
	FlagLazyInit              uint32 = 0x40
	FlagTwoLevel              uint32 = 0x80
	FlagForceFlat             uint32 = 0x100
	FlagNoMultiDefs           uint32 = 0x200
	FlagNoFixPrebinding       uint32 = 0x400
	FlagPrebindable           uint32 = 0x800
	FlagAllModsBound          uint32 = 0x1000
	FlagSubsectionsViaSymbols uint32 = 0x2000
	FlagCanonical             uint32 = 0x4000
	FlagWeakDefines           uint32 = 0x8000
	FlagBindsToWeak           uint32 = 0x10000
	FlagAllowStackExecution   uint32 = 0x20000
	FlagRootSafe              uint32 = 0x40000
	FlagSetuidSafe            uint32 = 0x80000
	FlagNoReexportedDylibs    uint32 = 0x100000
	FlagPIE                   uint32 = 0x200000
	FlagDeadStrippableDylib   uint32 = 0x400000
	FlagHasTLVDescriptors     uint32 = 0x800000
	FlagNoHeapExecution       uint32 = 0x1000000
	FlagAppExtensionSafe      uint32 = 0x2000000
)

// Section32は、32ビットMach-Oセクションヘッダーです。
type Section32 struct {
	Name     [16]byte
	Seg      [16]byte
	Addr     uint32
	Size     uint32
	Offset   uint32
	Align    uint32
	Reloff   uint32
	Nreloc   uint32
	Flags    uint32
	Reserve1 uint32
	Reserve2 uint32
}

// Section64は、64ビットMach-Oセクションヘッダーです。
type Section64 struct {
	Name     [16]byte
	Seg      [16]byte
	Addr     uint64
	Size     uint64
	Offset   uint32
	Align    uint32
	Reloff   uint32
	Nreloc   uint32
	Flags    uint32
	Reserve1 uint32
	Reserve2 uint32
	Reserve3 uint32
}

// Nlist32は、Mach-O 32ビットのシンボルテーブルエントリです。
type Nlist32 struct {
	Name  uint32
	Type  uint8
	Sect  uint8
	Desc  uint16
	Value uint32
}

// Nlist64は、Mach-O 64ビットのシンボルテーブルエントリです。
type Nlist64 struct {
	Name  uint32
	Type  uint8
	Sect  uint8
	Desc  uint16
	Value uint64
}

// Regs386は、Mach-O 386のレジスタ構造体です。
type Regs386 struct {
	AX    uint32
	BX    uint32
	CX    uint32
	DX    uint32
	DI    uint32
	SI    uint32
	BP    uint32
	SP    uint32
	SS    uint32
	FLAGS uint32
	IP    uint32
	CS    uint32
	DS    uint32
	ES    uint32
	FS    uint32
	GS    uint32
}

// RegsAMD64は、Mach-O AMD64のレジスタ構造体です。
type RegsAMD64 struct {
	AX    uint64
	BX    uint64
	CX    uint64
	DX    uint64
	DI    uint64
	SI    uint64
	BP    uint64
	SP    uint64
	R8    uint64
	R9    uint64
	R10   uint64
	R11   uint64
	R12   uint64
	R13   uint64
	R14   uint64
	R15   uint64
	IP    uint64
	FLAGS uint64
	CS    uint64
	FS    uint64
	GS    uint64
}
