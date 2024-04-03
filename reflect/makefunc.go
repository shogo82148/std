// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// MakeFuncの実装。

package reflect

<<<<<<< HEAD
// MakeFuncは、与えられたTypeの新しい関数を返します。
// この新しい関数は、呼び出されると以下の操作を実行します：
=======
// MakeFunc returns a new function of the given [Type]
// that wraps the function fn. When called, that new function
// does the following:
>>>>>>> upstream/master
//
//   - 引数をValuesのスライスに変換します。
//   - results := fn(args)を実行します。
//   - 結果を一つずつフォーマルな結果に対応するValuesのスライスとして返します。
//
<<<<<<< HEAD
// fnの実装は、引数のValueのスライスがtypで与えられた引数の数と型を持っていると仮定できます。
// typが可変長引数の関数を記述している場合、最後のValue自体が可変長引数を表すスライス（可変長引数の関数の本体と同じように）です。
// fnによって返される結果のValueのスライスは、typで与えられた結果の数と型を持つ必要があります。
//
// Value.Callメソッドを使用することで、呼び出し元はValuesを用いた型指定の関数を呼び出すことができます。
// 対照的に、MakeFuncは型指定の関数をValuesを用いて実装することを呼び出し元に許可します。
=======
// The implementation fn can assume that the argument [Value] slice
// has the number and type of arguments given by typ.
// If typ describes a variadic function, the final Value is itself
// a slice representing the variadic arguments, as in the
// body of a variadic function. The result Value slice returned by fn
// must have the number and type of results given by typ.
//
// The [Value.Call] method allows the caller to invoke a typed function
// in terms of Values; in contrast, MakeFunc allows the caller to implement
// a typed function in terms of Values.
>>>>>>> upstream/master
//
// ドキュメントのExamplesセクションには、さまざまな型のスワップ関数を構築する方法の説明が含まれています。
func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
