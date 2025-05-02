// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cpu implements processor feature detection
// used by the Go standard library.
package cpu

// DebugOptions is set to true by the runtime if the OS supports reading
// GODEBUG early in runtime startup.
// This should not be changed after it is initialized.
var DebugOptions bool

// CacheLinePad is used to pad structs to avoid false sharing.
type CacheLinePad struct{ _ [CacheLinePadSize]byte }

// CacheLineSize is the CPU's assumed cache line size.
// There is currently no runtime detection of the real cache line size
// so we use the constant per GOARCH CacheLinePadSize as an approximation.
var CacheLineSize uintptr = CacheLinePadSize

// The booleans in X86 contain the correspondingly named cpuid feature bit.
// HasAVX and HasAVX2 are only set if the OS does support XMM and YMM registers
// in addition to the cpuid feature bit being set.
// The struct is padded to avoid false sharing.
var X86 struct {
	_            CacheLinePad
	HasAES       bool
	HasADX       bool
	HasAVX       bool
	HasAVX2      bool
	HasAVX512F   bool
	HasAVX512BW  bool
	HasAVX512VL  bool
	HasBMI1      bool
	HasBMI2      bool
	HasERMS      bool
	HasFSRM      bool
	HasFMA       bool
	HasOSXSAVE   bool
	HasPCLMULQDQ bool
	HasPOPCNT    bool
	HasRDTSCP    bool
	HasSHA       bool
	HasSSE3      bool
	HasSSSE3     bool
	HasSSE41     bool
	HasSSE42     bool
	_            CacheLinePad
}

// The booleans in ARM contain the correspondingly named cpu feature bit.
// The struct is padded to avoid false sharing.
var ARM struct {
	_            CacheLinePad
	HasVFPv4     bool
	HasIDIVA     bool
	HasV7Atomics bool
	_            CacheLinePad
}

// The booleans in ARM64 contain the correspondingly named cpu feature bit.
// The struct is padded to avoid false sharing.
var ARM64 struct {
	_          CacheLinePad
	HasAES     bool
	HasPMULL   bool
	HasSHA1    bool
	HasSHA2    bool
	HasSHA512  bool
	HasCRC32   bool
	HasATOMICS bool
	HasCPUID   bool
	HasDIT     bool
	IsNeoverse bool
	_          CacheLinePad
}

// The booleans in Loong64 contain the correspondingly named cpu feature bit.
// The struct is padded to avoid false sharing.
var Loong64 struct {
	_         CacheLinePad
	HasLSX    bool
	HasLASX   bool
	HasCRC32  bool
	HasLAMCAS bool
	HasLAM_BH bool
	_         CacheLinePad
}

var MIPS64X struct {
	_      CacheLinePad
	HasMSA bool
	_      CacheLinePad
}

// For ppc64(le), it is safe to check only for ISA level starting on ISA v3.00,
// since there are no optional categories. There are some exceptions that also
// require kernel support to work (darn, scv), so there are feature bits for
// those as well. The minimum processor requirement is POWER8 (ISA 2.07).
// The struct is padded to avoid false sharing.
var PPC64 struct {
	_         CacheLinePad
	HasDARN   bool
	HasSCV    bool
	IsPOWER8  bool
	IsPOWER9  bool
	IsPOWER10 bool
	_         CacheLinePad
}

var S390X struct {
	_         CacheLinePad
	HasZARCH  bool
	HasSTFLE  bool
	HasLDISP  bool
	HasEIMM   bool
	HasDFP    bool
	HasETF3EH bool
	HasMSA    bool
	HasAES    bool
	HasAESCBC bool
	HasAESCTR bool
	HasAESGCM bool
	HasGHASH  bool
	HasSHA1   bool
	HasSHA256 bool
	HasSHA512 bool
	HasSHA3   bool
	HasVX     bool
	HasVXE    bool
	HasKDSA   bool
	HasECDSA  bool
	HasEDDSA  bool
	_         CacheLinePad
}

// RISCV64 contains the supported CPU features and performance characteristics for riscv64
// platforms. The booleans in RISCV64, with the exception of HasFastMisaligned, indicate
// the presence of RISC-V extensions.
// The struct is padded to avoid false sharing.
var RISCV64 struct {
	_                 CacheLinePad
	HasFastMisaligned bool
	HasV              bool
	HasZbb            bool
	_                 CacheLinePad
}

// Initialize examines the processor and sets the relevant variables above.
// This is called by the runtime package early in program initialization,
// before normal init functions are run. env is set by runtime if the OS supports
// cpu feature options in GODEBUG.
func Initialize(env string)
