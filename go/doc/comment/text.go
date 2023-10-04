// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comment

// Text returns a textual formatting of the Doc.
// See the [Printer] documentation for ways to customize the text output.
func (p *Printer) Text(d *Doc) []byte
