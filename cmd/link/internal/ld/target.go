// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"github.com/shogo82148/std/cmd/internal/objabi"
	"github.com/shogo82148/std/cmd/internal/sys"
)

// Target holds the configuration we're building for.
type Target struct {
	Arch *sys.Arch

	HeadType objabi.HeadType

	LinkMode  LinkMode
	BuildMode BuildMode

	linkShared    bool
	canUsePlugins bool
	IsELF         bool
}

func (t *Target) IsExe() bool

func (t *Target) IsShared() bool

func (t *Target) IsPlugin() bool

func (t *Target) IsInternal() bool

func (t *Target) IsExternal() bool

func (t *Target) IsPIE() bool

func (t *Target) IsSharedGoLink() bool

func (t *Target) CanUsePlugins() bool

func (t *Target) IsElf() bool

func (t *Target) IsDynlinkingGo() bool

// UseRelro reports whether to make use of "read only relocations" aka
// relro.
func (t *Target) UseRelro() bool

func (t *Target) Is386() bool

func (t *Target) IsARM() bool

func (t *Target) IsARM64() bool

func (t *Target) IsAMD64() bool

func (t *Target) IsMIPS() bool

func (t *Target) IsMIPS64() bool

func (t *Target) IsLOONG64() bool

func (t *Target) IsPPC64() bool

func (t *Target) IsRISCV64() bool

func (t *Target) IsS390X() bool

func (t *Target) IsWasm() bool

func (t *Target) IsLinux() bool

func (t *Target) IsDarwin() bool

func (t *Target) IsWindows() bool

func (t *Target) IsPlan9() bool

func (t *Target) IsAIX() bool

func (t *Target) IsSolaris() bool

func (t *Target) IsNetbsd() bool

func (t *Target) IsOpenbsd() bool

func (t *Target) IsFreebsd() bool

func (t *Target) IsBigEndian() bool

func (t *Target) UsesLibc() bool
