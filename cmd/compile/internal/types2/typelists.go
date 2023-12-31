// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

// TypeParamList holds a list of type parameters.
type TypeParamList struct{ tparams []*TypeParam }

// Len returns the number of type parameters in the list.
// It is safe to call on a nil receiver.
func (l *TypeParamList) Len() int

// At returns the i'th type parameter in the list.
func (l *TypeParamList) At(i int) *TypeParam

// TypeList holds a list of types.
type TypeList struct{ types []Type }

// Len returns the number of types in the list.
// It is safe to call on a nil receiver.
func (l *TypeList) Len() int

// At returns the i'th type in the list.
func (l *TypeList) At(i int) Type
