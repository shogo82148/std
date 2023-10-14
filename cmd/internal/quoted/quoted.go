// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package quoted provides string manipulation utilities.
package quoted

import (
	"github.com/shogo82148/std/flag"
)

// Split splits s into a list of fields,
// allowing single or double quotes around elements.
// There is no unescaping or other processing within
// quoted fields.
//
// Keep in sync with cmd/dist/quoted.go
func Split(s string) ([]string, error)

// Join joins a list of arguments into a string that can be parsed
// with Split. Arguments are quoted only if necessary; arguments
// without spaces or quotes are kept as-is. No argument may contain both
// single and double quotes.
func Join(args []string) (string, error)

// A Flag parses a list of string arguments encoded with Join.
// It is useful for flags like cmd/link's -extldflags.
type Flag []string

var _ flag.Value = (*Flag)(nil)

func (f *Flag) Set(v string) error

func (f *Flag) String() string
