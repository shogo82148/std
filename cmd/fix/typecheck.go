// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type TypeConfig struct {
	Type map[string]*Type
	Var  map[string]string
	Func map[string]string

	// 名前から型への外部マップ。
	// これはGoのソース自体には存在しない追加の型を提供します。
	// 現時点では、cgoによって生成された追加の型のみがあります。
	External map[string]string
}

// Typeは型のフィールドとメソッドを表します。
// もしフィールドやメソッドがここに見つからない場合は、次にEmbedリストで探します。
type Type struct {
	Field  map[string]string
	Method map[string]string
	Embed  []string
	Def    string
}
