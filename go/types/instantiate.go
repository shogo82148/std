// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements instantiation of generic types
// through substitution of type parameters by type arguments.

package types

// Instantiate instantiates the type orig with the given type arguments targs.
// orig must be a *Named or a *Signature type. If there is no error, the
// resulting Type is an instantiated type of the same kind (either a *Named or
// a *Signature). Methods attached to a *Named type are also instantiated, and
// associated with a new *Func that has the same position as the original
// method, but nil function scope.
//
// If ctxt is non-nil, it may be used to de-duplicate the instance against
// previous instances with the same identity. As a special case, generic
// *Signature origin types are only considered identical if they are pointer
// equivalent, so that instantiating distinct (but possibly identical)
// signatures will yield different instances. The use of a shared context does
// not guarantee that identical instances are deduplicated in all cases.
//
// If validate is set, Instantiate verifies that the number of type arguments
// and parameters match, and that the type arguments satisfy their
// corresponding type constraints. If verification fails, the resulting error
// may wrap an *ArgumentError indicating which type argument did not satisfy
// its corresponding type parameter constraint, and why.
//
// If validate is not set, Instantiate does not verify the type argument count
// or whether the type arguments satisfy their constraints. Instantiate is
// guaranteed to not return an error, but may panic. Specifically, for
// *Signature types, Instantiate will panic immediately if the type argument
// count is incorrect; for *Named types, a panic may occur later inside the
// *Named API.
func Instantiate(ctxt *Context, orig Type, targs []Type, validate bool) (Type, error)
