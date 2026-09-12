// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package variants

type Key struct {
	Arch string
	Size int
}

type Variant struct {
	Suffix          string
	Emulated        map[string]bool
	DefaultRequires string
}

// Name returns the variant name of type t.
// If receiver v is nil, returns t.
func (v *Variant) Name(t string) string

var Variants map[Key]*Variant
