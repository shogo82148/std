// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package load

import (
	"github.com/shogo82148/std/encoding/json"
	"github.com/shogo82148/std/io"
)

// A Printer reports output about a Package.
type Printer interface {
	Printf(pkg *Package, format string, args ...any)

	Errorf(pkg *Package, format string, args ...any)
}

// DefaultPrinter returns the default Printer.
func DefaultPrinter() Printer

// A TextPrinter emits text format output to Writer.
type TextPrinter struct {
	Writer io.Writer
}

func (p *TextPrinter) Printf(_ *Package, format string, args ...any)

func (p *TextPrinter) Errorf(_ *Package, format string, args ...any)

// A JSONPrinter emits output about a build in JSON format.
type JSONPrinter struct {
	enc *json.Encoder
}

func NewJSONPrinter(w io.Writer) *JSONPrinter

func (p *JSONPrinter) Printf(pkg *Package, format string, args ...any)

func (p *JSONPrinter) Errorf(pkg *Package, format string, args ...any)
