// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssahtml

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"

	"github.com/shogo82148/std/cmd/compile/internal/ssa"
)

type HTMLWriter struct {
	w              io.WriteCloser
	Func           *ssa.Func
	path           string
	dot            *dotWriter
	prevHash       []byte
	pendingPhases  []string
	pendingTitles  []string
	debugInfo      []string
	timeFormatting time.Duration
}

func NewHTMLWriter(path string, f *ssa.Func, cfgMask string, passes []ssa.Pass) *HTMLWriter

func (w *HTMLWriter) Enabled() bool

// Fatalf reports an error and exits.
func (w *HTMLWriter) Fatalf(msg string, args ...any)

// Logf calls the (w *HTMLWriter).Func's Logf method passing along a msg and args.
func (w *HTMLWriter) Logf(msg string, args ...any)

func (w *HTMLWriter) Close()

// WritePhase writes f in a column headed by title.
// phase is used for collapsing columns and should be unique across the table.
func (w *HTMLWriter) WritePhase(phase, title string)

// FatalCleanup should be called to do cleanup if the compilation is exiting early due to
// a fatal error.
func (w *HTMLWriter) FatalCleanup()

// FlushPhases collects any pending phases and titles, writes them to the html, and resets the pending slices.
func (w *HTMLWriter) FlushPhases()

// FuncLines contains source code for a function to be displayed
// in sources column.
type FuncLines struct {
	Filename    string
	StartLineno uint
	Lines       []string
}

// ByTopoCmp sorts topologically: target function is on top,
// followed by inlined functions sorted by filename and line numbers.
func ByTopoCmp(a, b *FuncLines) int

// WriteSources writes lines as source code in a column headed by title.
// phase is used for collapsing columns and should be unique across the table.
func (w *HTMLWriter) WriteSources(phase string, all []*FuncLines)

func (w *HTMLWriter) WriteAST(phase string, buf *bytes.Buffer)

// WriteColumn writes raw HTML in a column headed by title.
// It is intended for pre- and post-compilation log output.
func (w *HTMLWriter) WriteColumn(phase, title, class, html string)

func (w *HTMLWriter) WriteMultiTitleColumn(phase string, titles []string, class, html string)

func (w *HTMLWriter) Printf(msg string, v ...any)

func (w *HTMLWriter) WriteString(s string)

func (w *HTMLWriter) DebugInfo(format func(*ssa.Value) string)

func (w *HTMLWriter) TimeFormatting() time.Duration

func HTML(f *ssa.Func, phase string, dot *dotWriter, debugStr []string) string
