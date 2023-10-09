// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc

type Filter func(string) bool

// Filterは、フィルターfを通過しない名前のドキュメントを除外します。
// TODO（gri）："Type.Method"を名前として認識する。
func (p *Package) Filter(f Filter)
