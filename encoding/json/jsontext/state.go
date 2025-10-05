// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/iter"
)

// ErrDuplicateName indicates that a JSON token could not be
// encoded or decoded because it results in a duplicate JSON object name.
// This error is directly wrapped within a [SyntacticError] when produced.
//
// The name of a duplicate JSON object member can be extracted as:
//
//	err := ...
//	var serr jsontext.SyntacticError
//	if errors.As(err, &serr) && serr.Err == jsontext.ErrDuplicateName {
//		ptr := serr.JSONPointer // JSON pointer to duplicate name
//		name := ptr.LastToken() // duplicate name itself
//		...
//	}
//
// This error is only returned if [AllowDuplicateNames] is false.
var ErrDuplicateName = errors.New("duplicate object member name")

// ErrNonStringName indicates that a JSON token could not be
// encoded or decoded because it is not a string,
// as required for JSON object names according to RFC 8259, section 4.
// This error is directly wrapped within a [SyntacticError] when produced.
var ErrNonStringName = errors.New("object member name must be a string")

// Pointer is a JSON Pointer (RFC 6901) that references a particular JSON value
// relative to the root of the top-level JSON value.
//
// A Pointer is a slash-separated list of tokens, where each token is
// either a JSON object name or an index to a JSON array element
// encoded as a base-10 integer value.
// It is impossible to distinguish between an array index and an object name
// (that happens to be an base-10 encoded integer) without also knowing
// the structure of the top-level JSON value that the pointer refers to.
//
// There is exactly one representation of a pointer to a particular value,
// so comparability of Pointer values is equivalent to checking whether
// they both point to the exact same value.
type Pointer string

// IsValid reports whether p is a valid JSON Pointer according to RFC 6901.
// Note that the concatenation of two valid pointers produces a valid pointer.
func (p Pointer) IsValid() bool

// Contains reports whether the JSON value that p points to
// is equal to or contains the JSON value that pc points to.
func (p Pointer) Contains(pc Pointer) bool

// Parent strips off the last token and returns the remaining pointer.
// The parent of an empty p is an empty string.
func (p Pointer) Parent() Pointer

// LastToken returns the last token in the pointer.
// The last token of an empty p is an empty string.
func (p Pointer) LastToken() string

// AppendToken appends a token to the end of p and returns the full pointer.
func (p Pointer) AppendToken(tok string) Pointer

// Tokens returns an iterator over the reference tokens in the JSON pointer,
// starting from the first token until the last token (unless stopped early).
func (p Pointer) Tokens() iter.Seq[string]
