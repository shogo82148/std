// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc

// Deprecated returns the deprecation notice, if any, in the doc comment text doc.
// A deprecation notice is the text of a paragraph that starts with "Deprecated: ",
// stripped of the "Deprecated: " prefix.
// If there is no deprecation notice, it returns "", false.
//
// The doc argument is comment text as returned by [go/ast.CommentGroup.Text],
// such as [Func.Doc], [Package.Doc], [Type.Doc], or [Value.Doc].
func Deprecated(doc string) (text string, ok bool)
