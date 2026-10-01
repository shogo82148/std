// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specgen

// Index provides lookup of concrete spec functions by (receiver, name).
type Index struct {
	m     map[indexKey]*Func
	funcs []*Func
}

// NewIndex builds an Index over the given slice of spec functions.
func NewIndex(funcs []*Func) *Index

// Lookup finds the spec function for the given receiver and function/method name.
// For package-level functions, recv should be "".
func (idx *Index) Lookup(recv, name string) *Func

// Funcs returns all functions in the index.
func (idx *Index) Funcs() []*Func

// Len returns the number of functions in the index.
func (idx *Index) Len() int
