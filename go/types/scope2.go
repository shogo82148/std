// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements go/types-specific scope methods.
// These methods do not exist in types2.

package types

import "github.com/shogo82148/std/go/token"

// LookupParentは、sから始まるスコープの親チェーンをたどり、
// Lookup(name)がnilでないオブジェクトを返すスコープを見つけると、
// そのスコープとオブジェクトを返します。有効な位置posが指定された場合、
// pos以前に宣言されたオブジェクトのみが考慮されます。
// そのようなスコープとオブジェクトが存在しない場合、結果は(nil, nil)です。
// 結果が有効であることが保証されるのは、型チェックされたASTが
// 完全な位置情報を持っている場合のみです。
//
// オブジェクトがスコープに挿入されたときにすでに親を持っていた場合
// （Insertを参照）、obj.Parent()は返されたスコープと異なる場合があることに注意してください。
// これは、親がそれをエクスポートしたパッケージのスコープであるドットインポートされた
// オブジェクトでのみ発生します。
func (s *Scope) LookupParent(name string, pos token.Pos) (*Scope, Object)

// PosとEndは、スコープのソースコード範囲[pos, end)を表します。
// 結果が有効であることが保証されるのは、型チェックされたASTが
// 完全な位置情報を持っている場合のみです。範囲はUniverseスコープと
// パッケージスコープについては未定義です。
func (s *Scope) Pos() token.Pos
func (s *Scope) End() token.Pos

// Containsは、posがスコープの範囲内にあるかどうかを報告します。
// 結果が有効であることが保証されるのは、型チェックされたASTが
// 完全な位置情報を持っている場合のみです。
func (s *Scope) Contains(pos token.Pos) bool

// Innermostは、posを含む最も内側の（子の）スコープを返します。
// posがどのスコープ内にもない場合、結果はnilです。
// Universeスコープの場合も結果はnilです。
// 結果が有効であることが保証されるのは、型チェックされたASTが
// 完全な位置情報を持っている場合のみです。
func (s *Scope) Innermost(pos token.Pos) *Scope
