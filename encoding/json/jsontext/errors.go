// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

// SyntacticError is a description of a syntactic error that occurred when
// encoding or decoding JSON according to the grammar.
//
// The contents of this error as produced by this package may change over time.
type SyntacticError struct {
	requireKeyedLiterals
	nonComparable

	// ByteOffset indicates that an error occurred after this byte offset.
	ByteOffset int64
	// JSONPointer indicates that an error occurred within this JSON value
	// as indicated using the JSON Pointer notation (see RFC 6901).
	JSONPointer Pointer

	// Err is the underlying error.
	Err error
}

func (e *SyntacticError) Error() string

func (e *SyntacticError) Unwrap() error
