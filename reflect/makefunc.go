// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// MakeFuncの実装。

package reflect

// MakeFuncは、与えられた [Type] の新しい関数を返します。
// この新しい関数は、呼び出されると以下の操作を実行します：
//
//   - 引数をValuesのスライスに変換します。
//   - results := fn(args)を実行します。
//   - 結果を一つずつフォーマルな結果に対応するValuesのスライスとして返します。
//
// fnの実装は、引数の [Value] のスライスがtypで与えられた引数の数と型を持っていると仮定できます。
// typが可変長引数の関数を記述している場合、最後のValue自体が可変長引数を表すスライス（可変長引数の関数の本体と同じように）です。
// fnによって返される結果のValueのスライスは、typで与えられた結果の数と型を持つ必要があります。
//
// [Value.Call] メソッドを使用することで、呼び出し元はValuesを用いた型指定の関数を呼び出すことができます。
// 対照的に、MakeFuncは型指定の関数をValuesを用いて実装することを呼び出し元に許可します。
//
// ドキュメントのExamplesセクションには、さまざまな型のスワップ関数を構築する方法の説明が含まれています。
func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
