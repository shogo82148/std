// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// TypeはGoの型を表します。
// すべての型はTypeインターフェースを実装しています。
type Type interface {
	Underlying() Type

	String() string
}
