// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// Signatureは（ビルトインでない）関数またはメソッドの型を表します。
// シグネチャを同一性で比較する際は、レシーバは無視されます。
type Signature struct {

	// シグネチャのスコープを維持する必要があります（それを渡すのではなく、Funcオブジェクトに保存します）。
	// なぜなら、関数リテラルの型チェック時に一般的な型チェッカーが呼び出され、一般的なタイプが返されるからです。
	// 次に、*Signature を展開してリテラルの本体のためにスコープを使用します。
	rparams  *TypeParamList
	tparams  *TypeParamList
	scope    *Scope
	recv     *Var
	params   *Tuple
	results  *Tuple
	variadic bool
}

// NewSignatureは与えられたレシーバ、パラメータ、戻り値のための新しい関数型を返します。それぞれnilにすることもできます。もしvariadicがtrueに設定されている場合、関数は可変長引数を持ち、少なくとも1つのパラメータを持つ必要があります。また、最後のパラメータは無名のスライス型である必要があります。
//
// Deprecated: 代わりに型パラメータを許可する [NewSignatureType] を使用してください。
//
//go:fix inline
func NewSignature(recv *Var, params, results *Tuple, variadic bool) *Signature

// NewSignatureTypeは、与えられたレシーバ、レシーバ型パラメータ、型パラメータ、
// パラメータ、および結果のための新しい関数型を作成します。
// variadicが設定されている場合、paramsは少なくとも1つのパラメータを持つ必要があり、
// 最後のパラメータは無名のスライスまたは型セットが共通の基底型として
// 無名のスライスを持つ型パラメータである必要があります。
// 特別なケースとして、可変長引数シグネチャの場合、最後のパラメータは
// string型、または型セットにbyteスライスとstring型の組み合わせを含む
// 型パラメータであっても構いません。
// recvがnon-nilの場合、typeParamsは空である必要があります。recvTypeParamsが
// 空でない場合、recvはnon-nilである必要があります。
func NewSignatureType(recv *Var, recvTypeParams, typeParams []*TypeParam, params, results *Tuple, variadic bool) *Signature

// Recvはシグネチャsのレシーバー（メソッドの場合）を返します。関数の場合はnilを返します。
// 同一性を比較する際にレシーバーは無視されます。
//
// 抽象メソッドの場合、Recvは囲んでいるインターフェースを
// *[Named] または*[Interface] として返します。埋め込みにより、インターフェースには
// レシーバー型が異なるインターフェースであるメソッドが含まれる場合があります。
func (s *Signature) Recv() *Var

// TypeParamsはシグネチャsの型パラメータを返します。パラメータが存在しない場合はnilを返します。
func (s *Signature) TypeParams() *TypeParamList

// RecvTypeParams はシグネチャ s のレシーバー型パラメーターを返します。nil の場合もあります。
func (s *Signature) RecvTypeParams() *TypeParamList

// Paramsはシグネチャsのパラメータを返します。パラメータがない場合はnilを返します。
func (s *Signature) Params() *Tuple

// Resultsはシグネチャsの結果、またはnilを返します。
func (s *Signature) Results() *Tuple

// Variadicは、シグネチャsが可変長引数であるかどうかを報告します。
func (s *Signature) Variadic() bool

func (t *Signature) Underlying() Type
func (t *Signature) String() string
