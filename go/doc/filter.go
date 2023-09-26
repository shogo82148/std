// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc

type Filter func(string) bool

// Filter eliminates documentation for names that don't pass through the filter f.
// TODO: Recognize "Type.Method" as a name.
func (p *Package) Filter(f Filter)
