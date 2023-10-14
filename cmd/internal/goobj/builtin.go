// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goobj

// NBuiltin returns the number of listed builtin
// symbols.
func NBuiltin() int

// BuiltinName returns the name and ABI of the i-th
// builtin symbol.
func BuiltinName(i int) (string, int)

// BuiltinIdx returns the index of the builtin with the
// given name and abi, or -1 if it is not a builtin.
func BuiltinIdx(name string, abi int) int
