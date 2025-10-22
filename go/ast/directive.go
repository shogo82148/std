// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

import (
	"github.com/shogo82148/std/go/token"
)

// A Directive is a comment of this form:
//
//	//tool:name args
//
// For example, this directive:
//
//	//go:generate stringer -type Op -trimprefix Op
//
// would have Tool "go", Name "generate", and Args "stringer -type Op
// -trimprefix Op".
//
// While Args does not have a strict syntax, by convention it is a
// space-separated sequence of unquoted words, '"'-quoted Go strings, or
// '`'-quoted raw strings.
//
// See https://go.dev/doc/comment#directives for specification.
type Directive struct {
	Tool string
	Name string
	Args string

	// Slash is the position of the "//" at the beginning of the directive.
	Slash token.Pos

	// ArgsPos is the position where Args begins, based on the position passed
	// to ParseDirective.
	ArgsPos token.Pos
}

// ParseDirective parses a single comment line for a directive comment.
//
// If the line is not a directive comment, it returns false.
//
// The provided text must be a single line and should include the leading "//".
// If the text does not start with "//", it returns false.
//
// The caller may provide a file position of the start of c. This will be used
// to track the position of the arguments. This may be [Comment.Slash],
// synthesized by the caller, or simply 0. If the caller passes 0, then the
// positions are effectively byte offsets into the string c.
func ParseDirective(pos token.Pos, c string) (Directive, bool)

func (d *Directive) Pos() token.Pos
func (d *Directive) End() token.Pos

// A DirectiveArg is an argument to a directive comment.
type DirectiveArg struct {
	// Arg is the parsed argument string. If the argument was a quoted string,
	// this is its unquoted form.
	Arg string
	// Pos is the position of the first character in this argument.
	Pos token.Pos
}

// ParseArgs parses a [Directive]'s arguments using the standard convention,
// which is a sequence of tokens, where each token may be a bare word, or a
// double quoted Go string, or a back quoted raw Go string. Each token must be
// separated by one or more Unicode spaces.
//
// If the arguments do not conform to this syntax, it returns an error.
func (d *Directive) ParseArgs() ([]DirectiveArg, error)
