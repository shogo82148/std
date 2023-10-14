// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/io"
)

type HTMLWriter struct {
	w             io.WriteCloser
	Func          *Func
	path          string
	dot           *dotWriter
	prevHash      []byte
	pendingPhases []string
	pendingTitles []string
}

func NewHTMLWriter(path string, f *Func, cfgMask string) *HTMLWriter

// Fatalf reports an error and exits.
func (w *HTMLWriter) Fatalf(msg string, args ...interface{})

// Logf calls the (w *HTMLWriter).Func's Logf method passing along a msg and args.
func (w *HTMLWriter) Logf(msg string, args ...interface{})

func (w *HTMLWriter) Close()

// WritePhase writes f in a column headed by title.
// phase is used for collapsing columns and should be unique across the table.
func (w *HTMLWriter) WritePhase(phase, title string)

// FuncLines contains source code for a function to be displayed
// in sources column.
type FuncLines struct {
	Filename    string
	StartLineno uint
	Lines       []string
}

// ByTopo sorts topologically: target function is on top,
// followed by inlined functions sorted by filename and line numbers.
type ByTopo []*FuncLines

func (x ByTopo) Len() int
func (x ByTopo) Swap(i, j int)
func (x ByTopo) Less(i, j int) bool

// WriteSources writes lines as source code in a column headed by title.
// phase is used for collapsing columns and should be unique across the table.
func (w *HTMLWriter) WriteSources(phase string, all []*FuncLines)

func (w *HTMLWriter) WriteAST(phase string, buf *bytes.Buffer)

// WriteColumn writes raw HTML in a column headed by title.
// It is intended for pre- and post-compilation log output.
func (w *HTMLWriter) WriteColumn(phase, title, class, html string)

func (w *HTMLWriter) WriteMultiTitleColumn(phase string, titles []string, class, html string)

func (w *HTMLWriter) Printf(msg string, v ...interface{})

func (w *HTMLWriter) WriteString(s string)

func (v *Value) HTML() string

func (v *Value) LongHTML() string

func (b *Block) HTML() string

func (b *Block) LongHTML() string

func (f *Func) HTML(phase string, dot *dotWriter) string
