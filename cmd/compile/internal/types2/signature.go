// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

// A Signature represents a (non-builtin) function or method type.
// The receiver is ignored when comparing signatures for identity.
type Signature struct {
	// We need to keep the scope in Signature (rather than passing it around
	// and store it in the Func Object) because when type-checking a function
	// literal we call the general type checker which returns a general Type.
	// We then unpack the *Signature and use the scope for the literal body.
	rparams  *TypeParamList
	tparams  *TypeParamList
	scope    *Scope
	recv     *Var
	params   *Tuple
	results  *Tuple
	variadic bool
}

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
func NewSignatureType(recv *Var, recvTypeParams, typeParams []*TypeParam, params, results *Tuple, variadic bool) *Signature

// Recv returns the receiver of signature s (if a method), or nil if a
// function. It is ignored when comparing signatures for identity.
//
// For an abstract method, Recv returns the enclosing interface either
// as a *Named or an *Interface. Due to embedding, an interface may
// contain methods whose receiver type is a different interface.
func (s *Signature) Recv() *Var

// TypeParams returns the type parameters of signature s, or nil.
func (s *Signature) TypeParams() *TypeParamList

// RecvTypeParams returns the receiver type parameters of signature s, or nil.
func (s *Signature) RecvTypeParams() *TypeParamList

// Params returns the parameters of signature s, or nil.
func (s *Signature) Params() *Tuple

// Results returns the results of signature s, or nil.
func (s *Signature) Results() *Tuple

// Variadic reports whether the signature s is variadic.
func (s *Signature) Variadic() bool

func (s *Signature) Underlying() Type
func (s *Signature) String() string
