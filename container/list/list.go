// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージlistは双方向連結リストを実装します。
//
// リストを反復するには（l は *List）:
//
//	for e := l.Front(); e != nil; e = e.Next() {
//		// e.Value を使って何らかの処理を行う
//	}
package list

// Elementは連結リストの要素です。
type Element struct {
	// 要素からなる双方向連結リストにおける次および前のポインタです。
	// 実装を簡単にするため、内部的にはリスト l はリングとして実装され、
	// &l.root は最後の要素（l.Back()）の次要素であると同時に、
	// 最初の要素（l.Front()）の前要素でもあります。
	next, prev *Element

	// この要素が属するリストです。
	list *List

	// この要素に格納される値です。
	Value any
}

// Nextは次のリスト要素を返します。なければnilを返します。
func (e *Element) Next() *Element

// Prevは前のリスト要素を返します。なければnilを返します。
func (e *Element) Prev() *Element

// Listは双方向連結リストを表します。
// Listのゼロ値は、すぐに使用可能な空のリストです。
type List struct {
	root Element
	len  int
}

// Initはリストlを初期化またはクリアします。
func (l *List) Init() *List

// Newは初期化済みのリストを返します。
func New() *List

// Lenはリストlの要素数を返します。
// 計算量は O(1) です。
func (l *List) Len() int

// Frontはリストlの先頭要素を返します。リストが空ならnilを返します。
func (l *List) Front() *Element

// Backはリストlの末尾要素を返します。リストが空ならnilを返します。
func (l *List) Back() *Element

// Removeは、eがリストlの要素であれば l から e を削除します。
// 要素の値 e.Value を返します。
// 要素はnilであってはなりません。
func (l *List) Remove(e *Element) any

// PushFrontは値vを持つ新しい要素eをリストlの先頭に挿入し、eを返します。
func (l *List) PushFront(v any) *Element

// PushBackは値vを持つ新しい要素eをリストlの末尾に挿入し、eを返します。
func (l *List) PushBack(v any) *Element

// InsertBeforeは値vを持つ新しい要素eを mark の直前に挿入し、eを返します。
// markが l の要素でない場合、リストは変更されません。
// markはnilであってはなりません。
func (l *List) InsertBefore(v any, mark *Element) *Element

// InsertAfterは値vを持つ新しい要素eを mark の直後に挿入し、eを返します。
// markが l の要素でない場合、リストは変更されません。
// markはnilであってはなりません。
func (l *List) InsertAfter(v any, mark *Element) *Element

// MoveToFrontは要素eをリストlの先頭へ移動します。
// eが l の要素でない場合、リストは変更されません。
// 要素はnilであってはなりません。
func (l *List) MoveToFront(e *Element)

// MoveToBackは要素eをリストlの末尾へ移動します。
// eが l の要素でない場合、リストは変更されません。
// 要素はnilであってはなりません。
func (l *List) MoveToBack(e *Element)

// MoveBeforeは要素eを mark の前の新しい位置へ移動します。
// e または mark が l の要素でない場合、または e == mark の場合、
// リストは変更されません。
// 要素と mark はnilであってはなりません。
func (l *List) MoveBefore(e, mark *Element)

// MoveAfterは要素eを mark の後の新しい位置へ移動します。
// e または mark が l の要素でない場合、または e == mark の場合、
// リストは変更されません。
// 要素と mark はnilであってはなりません。
func (l *List) MoveAfter(e, mark *Element)

// PushBackListは別のリストのコピーをリストlの末尾に挿入します。
// リスト l と other は同一でも構いません。nilであってはなりません。
func (l *List) PushBackList(other *List)

// PushFrontListは別のリストのコピーをリストlの先頭に挿入します。
// リスト l と other は同一でも構いません。nilであってはなりません。
func (l *List) PushFrontList(other *List)
