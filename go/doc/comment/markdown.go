// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comment

// An mdPrinter holds the state needed for printing a Doc as Markdown.

// Markdown returns a Markdown formatting of the Doc.
// See the [Printer] documentation for ways to customize the Markdown output.
func (p *Printer) Markdown(d *Doc) []byte
