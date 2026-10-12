// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// IsFlagOp reports if v is an OP with the flag type.
func (v *Value) IsFlagOp() bool

// HasFlagInput reports whether v has a flag value as any of its inputs.
func (v *Value) HasFlagInput() bool
