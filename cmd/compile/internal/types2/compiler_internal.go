// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
)

// RenameResult takes an array of (result) fields and an index, and if the indexed field
// does not have a name and if the result in the signature also does not have a name,
// then the signature and field are renamed to
//
//	fmt.Sprintf("#rv%d", i+1)
//
// the newly named object is inserted into the signature's scope,
// and the object and new field name are returned.
//
// The intended use for RenameResult is to allow rangefunc to assign results within a closure.
// This is a hack, as narrowly targeted as possible to discourage abuse.
func (s *Signature) RenameResult(results []*syntax.Field, i int) (*Var, *syntax.Name)
