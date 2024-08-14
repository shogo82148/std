// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.swissmap

package reflect

// MapOfは、指定されたキーと要素の型を持つマップ型を返します。
// 例えば、kがintを表し、eがstringを表す場合、
// MapOf(k, e)はmap[int]stringを表します。
//
// キーの型が有効なマップキーの型でない場合（つまり、Goの==演算子を
// 実装していない場合）、MapOfはパニックを起こします。
func MapOf(key, elem Type) Type

// MapIndexは、マップv内のキーに関連付けられた値を返します。
// vのKindが [Map] でない場合、パニックを起こします。
// キーがマップ内に見つからない場合、またはvがnilマップを表す場合、ゼロ値を返します。
// Goと同様に、キーの値はマップのキー型に代入可能でなければなりません。
func (v Value) MapIndex(key Value) Value

// MapKeysは、マップに存在するすべてのキーを含むスライスを返します。
// 順序は指定されていません。
// vのKindが [Map] でない場合、パニックを起こします。
// vがnilマップを表す場合、空のスライスを返します。
func (v Value) MapKeys() []Value

// MapIterは、マップを範囲指定して反復するためのイテレータです。
// 詳細は [Value.MapRange] を参照してください。
type MapIter struct {
	m     Value
	hiter hiter
}

// Keyは、iterの現在のマップエントリのキーを返します。
func (iter *MapIter) Key() Value

// SetIterKeyは、iterの現在のマップエントリのキーをvに割り当てます。
// これはv.Set(iter.Key())と同等ですが、新しいValueを割り当てることを避けます。
// Goと同様に、キーはvの型に代入可能でなければならず、
// 非公開フィールドから派生したものであってはなりません。
func (v Value) SetIterKey(iter *MapIter)

// Valueは、iterの現在のマップエントリの値を返します。
func (iter *MapIter) Value() Value

// SetIterValueは、iterの現在のマップエントリの値をvに割り当てます。
// これはv.Set(iter.Value())と同等ですが、新しいValueを割り当てることを避けます。
// Goと同様に、値はvの型に代入可能でなければならず、
// 非公開フィールドから派生したものであってはなりません。
func (v Value) SetIterValue(iter *MapIter)

// Nextはマップイテレータを進め、次のエントリがあるかどうかを報告します。
// iterが終了した場合、falseを返します。
// その後の [MapIter.Key]、[MapIter.Value]、または [MapIter.Next] の呼び出しはパニックを引き起こします。
func (iter *MapIter) Next() bool

// Resetは、iterをvを反復するように変更します。
// vのKindが [Map] でない場合、またはvがゼロ値でない場合、パニックを起こします。
// Reset(Value{})は、iterがどのマップも参照しないようにします。
// これにより、以前に反復されたマップがガベージコレクションされる可能性があります。
func (iter *MapIter) Reset(v Value)

// MapRangeはマップの範囲イテレータを返します。
// vのKindが [Map] でない場合、パニックを起こします。
//
// イテレータを進めるには [MapIter.Next] を呼び出し、各エントリにアクセスするには [MapIter.Key]/[MapIter.Value] を呼び出します。
// イテレータが終了した場合、[MapIter.Next] はfalseを返します。
// MapRangeはrangeステートメントと同じイテレーションセマンティクスに従います。
//
// 例:
//
//	iter := reflect.ValueOf(m).MapRange()
//	for iter.Next() {
//		k := iter.Key()
//		v := iter.Value()
//		...
//	}
func (v Value) MapRange() *MapIter

// SetMapIndexは、マップv内のキーに関連付けられた要素をelemに設定します。
// vのKindが [Map] でない場合、パニックを起こします。
// elemがゼロ値の場合、SetMapIndexはマップからキーを削除します。
// それ以外の場合、vがnilマップを保持している場合、SetMapIndexはパニックを起こします。
// Goと同様に、キーの要素はマップのキー型に代入可能でなければならず、
// 要素の値はマップの要素型に代入可能でなければなりません。
func (v Value) SetMapIndex(key, elem Value)
