// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージringは循環リストに対する操作を実装します。
package ring

// Ringは循環リスト、またはリングの要素です。
// リングには先頭も末尾もありません。どのリング要素へのポインタも
// リング全体への参照として機能します。空のリングはnilのRingポインタで
// 表されます。Ringのゼロ値は、Valueがnilの1要素のリングです。
type Ring struct {
	next, prev *Ring
	Value      any
}

// Nextは次のリング要素を返します。rは空であってはなりません。
func (r *Ring) Next() *Ring

// Prevは前のリング要素を返します。rは空であってはなりません。
func (r *Ring) Prev() *Ring

// Moveは、リング内を n % r.Len() 個ぶん、後方へ（n < 0）または前方へ
// （n >= 0）移動し、そのリング要素を返します。rは空であってはなりません。
func (r *Ring) Move(n int) *Ring

// Newはn個の要素からなるリングを作成します。
func New(n int) *Ring

// Linkはリングrとリングsを接続し、r.Next()がsになるようにして、
// r.Next()の元の値を返します。rは空であってはなりません。
//
// rとsが同じリングを指している場合、それらをリンクすると
// rとsの間の要素がリングから削除されます。削除された要素は
// サブリングを形成し、結果はそのサブリングへの参照になります
// （要素が削除されなかった場合でも、結果はr.Next()の元の値であり、
// nilではありません）。
//
// rとsが異なるリングを指している場合、それらをリンクすると
// sの要素がrの後ろに挿入された1つのリングが作成されます。
// 結果は、挿入後のsの最後の要素の次の要素を指します。
func (r *Ring) Link(s *Ring) *Ring

// Unlinkは、r.Next()から始まるリングrの n % r.Len() 個の要素を削除します。
// n % r.Len() == 0 の場合、rは変更されません。
// 結果は削除されたサブリングです。rは空であってはなりません。
func (r *Ring) Unlink(n int) *Ring

// Lenはリングr内の要素数を計算します。
// 実行時間は要素数に比例します。
func (r *Ring) Len() int

// Doはリングの各要素に対して、前方向の順序で関数fを呼び出します。
// fが*rを変更した場合、Doの振る舞いは未定義です。
func (r *Ring) Do(f func(any))
