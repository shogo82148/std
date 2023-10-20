// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comment

// HTML returns an HTML formatting of the [Doc].
// See the [Printer] documentation for ways to customize the HTML output.
func (p *Printer) HTML(d *Doc) []byte
