// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc

// A methodSet describes a set of methods. Entries where Decl == nil are conflict
// entries (more than one method with the same name at the same embedding level).

// An embeddedSet describes a set of embedded types.

// A namedType represents a named unqualified (package local, or possibly
// predeclared) type. The namedType for a type name is always found via
// reader.lookupType.

// reader accumulates documentation for a single package.
// It modifies the AST: Comments (declaration documentation)
// that have been collected by the reader are set to nil
// in the respective AST nodes so that they are not printed
// twice (once when printing the documentation and once when
// printing the corresponding AST node).

// IsPredeclared reports whether s is a predeclared identifier.
func IsPredeclared(s string) bool
