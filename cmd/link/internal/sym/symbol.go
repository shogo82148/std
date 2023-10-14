// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sym

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

const (
	SymVerABI0        = 0
	SymVerABIInternal = 1
	SymVerABICount    = 2
	SymVerStatic      = 10
)

func ABIToVersion(abi obj.ABI) int

func VersionToABI(v int) (obj.ABI, bool)
