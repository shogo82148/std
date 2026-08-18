// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// PhiElimValue tries to convert the phi v to a copy.
func PhiElimValue(v *Value) bool
