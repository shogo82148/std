// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.swissmap

package reflect

import (
	"github.com/shogo82148/std/internal/runtime/maps"
)

// MapOfは、指定されたキーと要素の型を持つマップ型を返します。
// 例えば、kがintを表し、eがstringを表す場合、
// MapOf(k, e)はmap[int]stringを表します。
//
// キーの型が有効なマップのキーの型でない場合（つまり、Goの==演算子を
// 実装していない場合）、MapOfはパニックします。
func MapOf(key, elem Type) Type

// MapIndexは、マップv内のkeyに関連付けられた値を返します。
// vのKindが [Map] でない場合はパニックします。
// keyがマップ内に見つからない場合、またはvがnilマップを表す場合は、ゼロ値のValueを返します。
// Goと同様に、keyの値はマップのキーの型に代入可能でなければなりません。
func (v Value) MapIndex(key Value) Value

// MapKeysは、マップに存在するすべてのキーを不特定の順序で含むスライスを返します。
// vのKindが [Map] でない場合はパニックします。
// vがnilマップを表す場合は空のスライスを返します。
func (v Value) MapKeys() []Value

// MapIterは、マップをrangeで走査するためのイテレーターです。
// [Value.MapRange] を参照してください。
type MapIter struct {
	m     Value
	hiter maps.Iter
}

// Keyは、iterの現在のマップエントリーのキーを返します。
func (iter *MapIter) Key() Value

// SetIterKeyは、iterの現在のマップエントリーのキーをvに代入します。
// これはv.Set(iter.Key())と同等ですが、新しいValueの割り当てを回避します。
// Goと同様に、キーはvの型に代入可能でなければならず、
// エクスポートされていないフィールドから派生したものであってはなりません。
// [Value.CanSet] がfalseを返す場合はパニックします。
func (v Value) SetIterKey(iter *MapIter)

// Valueは、iterの現在のマップエントリーの値を返します。
func (iter *MapIter) Value() Value

// SetIterValueは、iterの現在のマップエントリーの値をvに代入します。
// これはv.Set(iter.Value())と同等ですが、新しいValueの割り当てを回避します。
// Goと同様に、値はvの型に代入可能でなければならず、
// エクスポートされていないフィールドから派生したものであってはなりません。
// [Value.CanSet] がfalseを返す場合はパニックします。
func (v Value) SetIterValue(iter *MapIter)

// Nextはマップイテレーターを進め、別のエントリーがあるかどうかを報告します。
// iterが使い果たされた場合はfalseを返します。その後の
// [MapIter.Key]、[MapIter.Value]、[MapIter.Next] の呼び出しはパニックします。
func (iter *MapIter) Next() bool

// Resetは、vを走査するようにiterを変更します。
// vのKindが [Map] ではなく、vがゼロ値のValueでもない場合はパニックします。
// Reset(Value{})はiterがどのマップも参照しないようにし、
// 以前に走査していたマップがガベージコレクションされることを可能にします。
func (iter *MapIter) Reset(v Value)

// MapRangeは、マップのrangeイテレーターを返します。
// vのKindが [Map] でない場合はパニックします。
//
// イテレーターを進めるには [MapIter.Next] を呼び出し、各エントリーにアクセスするには
// [MapIter.Key]/[MapIter.Value] を呼び出します。
// イテレーターが使い果たされると、[MapIter.Next] はfalseを返します。
// MapRangeはrange文と同じ反復のセマンティクスに従います。
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

// SetMapIndexは、マップv内のkeyに関連付けられた要素をelemに設定します。
// vのKindが [Map] でない場合はパニックします。
// elemがゼロ値のValueの場合、SetMapIndexはマップからキーを削除します。
// それ以外の場合、vがnilマップを保持していると、SetMapIndexはパニックします。
// Goと同様に、keyの要素はマップのキーの型に代入可能でなければならず、
// elemの値はマップの要素の型に代入可能でなければなりません。
func (v Value) SetMapIndex(key, elem Value)
