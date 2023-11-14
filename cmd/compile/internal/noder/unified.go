// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package noder

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// LookupMethodFunc returns the ir.Func for an arbitrary full symbol name if
// that function exists in the set of available export data.
//
// This allows lookup of arbitrary functions and methods that aren't otherwise
// referenced by the local package and thus haven't been read yet.
//
// TODO(prattmic): Does not handle instantiation of generic types. Currently
// profiles don't contain the original type arguments, so we won't be able to
// create the runtime dictionaries.
//
// TODO(prattmic): Hit rate of this function is usually fairly low, and errors
// are only used when debug logging is enabled. Consider constructing cheaper
// errors by default.
func LookupFunc(fullName string) (*ir.Func, error)
