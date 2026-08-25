// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// ArchInfo contains all architecture-specific naming conventions.
type ArchInfo struct {
	Arch      string
	ArchUpper string
	ObjArch   string
	// GoTypeArch identifies a distinct simdgen target and is the single tag that
	// names its generated output: the Go API files (types_<GoTypeArch>.go,
	// ops_<GoTypeArch>.go) and, in the backend package, simdssa_<GoTypeArch>.go.
	// It equals Arch for amd64/arm64, but is "sve" for the SVE target, so SVE's
	// output sits alongside — not on top of — the NEON arm64 files even though
	// both compile as GOARCH arm64.
	//
	// It is also the key of the shared simdgenericOps.go merge: each target's
	// run tags its generic ops with // ARCH:<GoTypeArch> so a later run for a
	// different target unions its ops in without dropping the others. Keying on
	// Arch instead would make an SVE run strip the "arm64" tag off every NEON
	// generic op (and drop arm64-only ones), since SVE and NEON share Arch.
	//
	// TODO: once the NEON and SVE type/op sets are unified, arm64 and sve can
	// collapse back into a single target and this second tag can go away.
	GoTypeArch string
	// SIMDTag names this target's generated files and functions in the
	// architecture-agnostic backend directories: ssa/_gen/simd<TAG>ops.go and
	// simd<TAG>.rules, and ssagen/simd<TAG>intrinsics.go, along with the
	// simd<TAG>Ops and simd<TAG>Intrinsics functions they define.
	SIMDTag         string
	RegInfoKeys     []string
	RegInfoSet      map[string]bool
	RegInfoParams   string
	GeneratedHeader string
	// Scalable reports that this target's vectors have a length that is only
	// known at run time, so the generated types carry no fixed lane count. It is
	// a property of the instruction set, independent of whether the target
	// shares a backend package: a future RVV target is scalable and owns its
	// package, LSX/LASX are neither.
	Scalable     bool
	Arrangements []string
}

// GetArchInfo returns architecture-specific information based on the target architecture.
func GetArchInfo(arch string) (ArchInfo, error)

// CurrentArch returns the ArchInfo for the current FlagArch setting.
func CurrentArch() ArchInfo
