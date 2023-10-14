// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sym

import (
	"github.com/shogo82148/std/cmd/internal/objabi"
	"github.com/shogo82148/std/cmd/internal/sys"
)

// RelocVariant is a linker-internal variation on a relocation.
type RelocVariant uint8

const (
	RV_NONE RelocVariant = iota
	RV_POWER_LO
	RV_POWER_HI
	RV_POWER_HA
	RV_POWER_DS

	// RV_390_DBL is a s390x-specific relocation variant that indicates that
	// the value to be placed into the relocatable field should first be
	// divided by 2.
	RV_390_DBL

	RV_CHECK_OVERFLOW RelocVariant = 1 << 7
	RV_TYPE_MASK      RelocVariant = RV_CHECK_OVERFLOW - 1
)

func RelocName(arch *sys.Arch, r objabi.RelocType) string
