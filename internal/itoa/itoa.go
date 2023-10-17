// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple conversions to avoid depending on strconv.

package itoa

// Itoa converts val to a decimal string.
func Itoa(val int) string

// Uitoa converts val to a decimal string.
func Uitoa(val uint) string

// Uitox converts val (a uint) to a hexadecimal string.
func Uitox(val uint) string
