// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ppc64 || ppc64le
// +build ppc64 ppc64le

package runtime

// cpu can be tested at runtime in go assembler code to check for
// a certain ISA level or hardware capability, for example:
//	  ·cpu+facilities_hasVSX(SB) for checking the availability of VSX
//	  or
//	  ·cpu+facilities_isPOWER7(SB) for checking if the processor implements
//	  ISA 2.06 instructions.
