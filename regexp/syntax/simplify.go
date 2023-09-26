// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

// Simplify returns a regexp equivalent to re but without counted repetitions
// and with various other simplifications, such as rewriting /(?:a+)+/ to /a+/.
// The resulting regexp will execute correctly but its string representation
// will not produce the same parse tree, because capturing parentheses
// may have been duplicated or removed.  For example, the simplified form
// for /(x){1,2}/ is /(x)(x)?/ but both parentheses capture as $1.
// The returned regexp may share structure with or be the original.
func (re *Regexp) Simplify() *Regexp
