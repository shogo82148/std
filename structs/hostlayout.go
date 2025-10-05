// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package structs

// HostLayoutは構造体がホストのメモリレイアウトを使用することを示します。
// HostLayout型のフィールドを持つ構造体は、メモリ上でホストの期待に従って、一般的にホストのC ABIに従ってレイアウトされます。
//
// HostLayoutは、含まれる構造体の他の構造体型フィールド内のレイアウトや、HostLayoutでマークされた構造体を含む構造体のレイアウトには影響しません。
//
// 慣例として、HostLayoutは「_」という名前のフィールド型として、構造体定義の先頭に配置して使用します。
type HostLayout struct {
	_ hostLayout
}
