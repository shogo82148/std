// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specgen

// Index provides lookup of concrete spec functions by (receiver, name), and of
// API types by name.
type Index struct {
	m     map[indexKey]*Func
	funcs []*Func

	typeMap map[string]*TypeDecl
	types   []*TypeDecl
}

// NewIndex builds an Index over the given spec functions and API types.
func NewIndex(funcs []*Func, types []*TypeDecl) *Index

// LoadIndex loads a Go SIMD spec from the package in directory dir, like
// [Load], and returns an [Index] over it.
func LoadIndex(dir string, opts *LoadOptions) (*Index, error)

// Lookup finds the spec function for the given receiver and function/method name.
// For package-level functions, recv should be "".
func (idx *Index) Lookup(recv, name string) *Func

// IsExperiment reports whether the function or method recv.name is
// experimental, meaning it is enabled only by GOEXPERIMENT=simd rather than
// being part of the generally available API (enabled by
// GOEXPERIMENT=simdstable). An API is experimental if it isn't covered by spec
// or is marked //specgen:experiment. For package-level functions, recv should
// be "".
//
// Some APIs not covered by spec are generally available anyway, but that's up
// to the caller to decide.
func (idx *Index) IsExperiment(recv, name string) bool

// LookupType finds the API type with the given name, or returns nil.
func (idx *Index) LookupType(name string) *TypeDecl

// Funcs returns all functions in the index.
func (idx *Index) Funcs() []*Func

// Types returns all API types in the index.
func (idx *Index) Types() []*TypeDecl

// Len returns the number of functions in the index.
func (idx *Index) Len() int
