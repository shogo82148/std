// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"github.com/shogo82148/std/io"
)

type File struct {
	w      io.Writer
	funcs  []*Func
	consts []fileConst
}

func NewFile(w io.Writer) *File

func (f *File) AddFunc(fn *Func)

func (f *File) AddConst(name string, data any)

type Func struct {
	name  string
	nArgs int
	idGen int
	ops   []*op
}

func NewFunc(name string) *Func

func Arg[W wrap[T], T Word](fn *Func) T

func Return(results ...Value)

func (f *File) Compile()
