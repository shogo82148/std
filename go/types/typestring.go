// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements printing of types.

package types

import (
	"github.com/shogo82148/std/bytes"
)

// A Qualifier controls how named package-level objects are printed in
// calls to [TypeString], [ObjectString], and [SelectionString].
//
// These three formatting routines call the Qualifier for each
// package-level object O, and if the Qualifier returns a non-empty
// string p, the object is printed in the form p.O.
// If it returns an empty string, only the object name O is printed.
//
// Using a nil Qualifier is equivalent to using (*[Package]).Path: the
// object is qualified by the import path, e.g., "encoding/json.Marshal".
type Qualifier func(*Package) string

// RelativeTo returns a [Qualifier] that fully qualifies members of
// all packages other than pkg.
func RelativeTo(pkg *Package) Qualifier

// TypeString returns the string representation of typ.
// The [Qualifier] controls the printing of
// package-level objects, and may be nil.
func TypeString(typ Type, qf Qualifier) string

// WriteType writes the string representation of typ to buf.
// The [Qualifier] controls the printing of
// package-level objects, and may be nil.
func WriteType(buf *bytes.Buffer, typ Type, qf Qualifier)

// WriteSignature writes the representation of the signature sig to buf,
// without a leading "func" keyword. The [Qualifier] controls the printing
// of package-level objects, and may be nil.
func WriteSignature(buf *bytes.Buffer, sig *Signature, qf Qualifier)
