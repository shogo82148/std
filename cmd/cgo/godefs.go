// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// gofmtLineReplacer is used to put a gofmt-formatted string for an
// AST expression onto a single line. The lexer normally inserts a
// semicolon at each newline, so we can replace newline with semicolon.
// However, we can't do that in cases where the lexer would not insert
// a semicolon. We only have to worry about cases that can occur in an
// expression passed through gofmt, which means composite literals and
// (due to the printer possibly inserting newlines because of position
// information) operators.
