// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

// An Alias represents an alias type.
// Whether or not Alias types are created is controlled by the
// gotypesalias setting with the GODEBUG environment variable.
// For gotypesalias=1, alias declarations produce an Alias type.
// Otherwise, the alias information is only in the type name,
// which points directly to the actual (aliased) type.
type Alias struct {
	obj     *TypeName
	tparams *TypeParamList
	fromRHS Type
	actual  Type
}

// NewAlias creates a new Alias type with the given type name and rhs.
// rhs must not be nil.
func NewAlias(obj *TypeName, rhs Type) *Alias

func (a *Alias) Obj() *TypeName
func (a *Alias) Underlying() Type
func (a *Alias) String() string

// Unalias returns t if it is not an alias type;
// otherwise it follows t's alias chain until it
// reaches a non-alias type which is then returned.
// Consequently, the result is never an alias type.
func Unalias(t Type) Type
