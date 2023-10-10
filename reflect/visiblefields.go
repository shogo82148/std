// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

// VisibleFieldsはtの中のすべての可視フィールドを返します。tはstruct型である必要があります。
// フィールドは、FieldByName呼び出しで直接アクセス可能ならば、可視として定義されます。
// 返されるフィールドには、無名structメンバー内のフィールドと非公開フィールドが含まれます。
// これらは、struct内で見つかった順序と同じ並び順になります。無名フィールドは、即座にそれに続く昇格フィールドが続きます。
//
// 返されるスライスの各要素eに対応するフィールドは、値vのタイプtからv.FieldByIndex(e.Index)を呼び出すことで取得できます。
func VisibleFields(t Type) []StructField
