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
<<<<<<< HEAD
// Deprecated: 代わりに型パラメータを使用できる [NewSignatureType] を使用してください。
func NewSignature(recv *Var, params, results *Tuple, variadic bool) *Signature

// NewSignatureTypeは、与えられたレシーバ、レシーバタイプパラメータ、
// タイプパラメータ、パラメータ、および結果に対して新しい関数型を作成します。
// variadicが設定されている場合、paramsは少なくとも1つのパラメータを保持している必要があり、
// 最後のパラメータのコア型は未命名スライスまたはバイト文字列型である必要があります。
// recvがnilでない場合、typeParamsは空でなければなりません。
// recvTypeParamsが空でない場合、recvはnilではない必要があります。
=======
// Deprecated: Use [NewSignatureType] instead which allows for type parameters.
//
//go:fix inline
func NewSignature(recv *Var, params, results *Tuple, variadic bool) *Signature

// NewSignatureType creates a new function type for the given receiver,
// receiver type parameters, type parameters, parameters, and results.
// If variadic is set, params must hold at least one parameter and the
// last parameter must be an unnamed slice or a type parameter whose
// type set has an unnamed slice as common underlying type.
// As a special case, for variadic signatures the last parameter may
// also be a string type, or a type parameter containing a mix of byte
// slices and string types in its type set.
// If recv is non-nil, typeParams must be empty. If recvTypeParams is
// non-empty, recv must be non-nil.
>>>>>>> upstream/release-branch.go1.25
func NewSignatureType(recv *Var, recvTypeParams, typeParams []*TypeParam, params, results *Tuple, variadic bool) *Signature

// Recv returns the receiver of signature s (if a method), or nil if a
// function. It is ignored when comparing signatures for identity.
//
// For an abstract method, Recv returns the enclosing interface either
// as a *[Named] or an *[Interface]. Due to embedding, an interface may
// contain methods whose receiver type is a different interface.
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
