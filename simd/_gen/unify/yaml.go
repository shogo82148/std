// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unify

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
)

// ReadOpts provides options to [Read] and related functions. The zero value is
// the default options.
type ReadOpts struct {
	// FS, if non-nil, is the file system from which to resolve !import file
	// names.
	FS fs.FS
}

// Read reads a [Closure] in YAML format from r, using path for error messages.
//
// It maps YAML nodes into terminal Values as follows:
//
// - "_" or !top _ is the top value ([Top]).
//
// - "_|_" or !bottom _ is the bottom value. This is an error during
// unmarshaling, but can appear in marshaled values.
//
// - "$<name>" or !var <name> is a variable ([Var]). Everywhere the same name
// appears within a single unmarshal operation, it is mapped to the same
// variable. Different unmarshal operations get different variables, even if
// they have the same string name.
//
// - !regex "x" is a regular expression ([String]), as is any string that
// doesn't match "_", "_|_", or "$...". Regular expressions are implicitly
// anchored at the beginning and end. If the string doesn't contain any
// meta-characters (that is, it's a "literal" regular expression), then it's
// treated as an exact string.
//
// - !string "x", or any int, float, bool, or binary value is an exact string
// ([String]).
//
// - !regex [x, y, ...] is an intersection of regular expressions ([String]).
//
// It maps YAML nodes into non-terminal Values as follows:
//
// - Sequence nodes like [x, y, z] are tuples ([Tuple]).
//
// - !repeat [x] is a repeated tuple ([Tuple]), which is 0 or more instances of
// x. There must be exactly one element in the list.
//
// - Mapping nodes like {a: x, b: y} are defs ([Def]). Any fields not listed are
// implicitly top.
//
// - !sum [x, y, z] is a sum of its children. This can be thought of as a union
// of the values x, y, and z, or as a non-deterministic choice between x, y, and
// z. If a variable appears both inside the sum and outside of it, only the
// non-deterministic choice view really works. The unifier does not directly
// implement sums; instead, this is decoded as a fresh variable that's
// simultaneously bound to x, y, and z.
//
// - !import glob is like a !sum, but its children are read from all files
// matching the given glob pattern, which is interpreted relative to the current
// file path. Each file gets its own variable scope.
func Read(r io.Reader, path string, opts ReadOpts) (Closure, error)

// ReadFile reads a [Closure] in YAML format from a file.
//
// The file must consist of a single YAML document.
//
// If opts.FS is not set, this sets it to a FS rooted at path's directory.
//
// See [Read] for details.
func ReadFile(path string, opts ReadOpts) (Closure, error)

// UnmarshalYAML implements [yaml.Unmarshaler].
//
// Since there is no way to pass [ReadOpts] to this function, it assumes default
// options.
func (c *Closure) UnmarshalYAML(node *yaml.Node) error

func (c Closure) MarshalYAML() (any, error)

func (c Closure) String() string

func (v *Value) MarshalYAML() (any, error)

func (v *Value) String() string
