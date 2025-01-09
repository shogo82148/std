// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fips140

var Enabled bool

// Supported returns an error if FIPS 140-3 mode can't be enabled.
func Supported() error

func Name() string

func Version() string
