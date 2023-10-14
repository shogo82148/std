// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lazyregexp is a thin wrapper over regexp, allowing the use of global
// regexp variables without forcing them to be compiled at init.
package lazyregexp

import (
	"github.com/shogo82148/std/regexp"
	"github.com/shogo82148/std/sync"
)

// Regexp is a wrapper around regexp.Regexp, where the underlying regexp will be
// compiled the first time it is needed.
type Regexp struct {
	str  string
	once sync.Once
	rx   *regexp.Regexp
}

func (r *Regexp) FindSubmatch(s []byte) [][]byte

func (r *Regexp) FindStringSubmatch(s string) []string

func (r *Regexp) FindStringSubmatchIndex(s string) []int

func (r *Regexp) ReplaceAllString(src, repl string) string

func (r *Regexp) FindString(s string) string

func (r *Regexp) FindAllString(s string, n int) []string

func (r *Regexp) MatchString(s string) bool

func (r *Regexp) SubexpNames() []string

// New creates a new lazy regexp, delaying the compiling work until it is first
// needed. If the code is being run as part of tests, the regexp compiling will
// happen immediately.
func New(str string) *Regexp
