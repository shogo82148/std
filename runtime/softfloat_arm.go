// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Software floating point interpretation of ARM 7500 FP instructions.
// The interpretation is not bit compatible with the 7500.
// It uses true little-endian doubles, while the 7500 used mixed-endian.

package runtime

// conditions array record the required CPSR cond field for the
// first 5 pairs of conditional execution opcodes
// higher 4 bits are must set, lower 4 bits are must clear
