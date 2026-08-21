// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

func (v *Value) HTML() string

func (v *Value) LongHTML(debugStr string) string

func (b *Block) HTML() string

func (b *Block) LongHTML() string

func (b *Block) UnlikelyIndex() int
