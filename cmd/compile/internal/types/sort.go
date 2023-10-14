// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// MethodsByName sorts methods by name.
type MethodsByName []*Field

func (x MethodsByName) Len() int
func (x MethodsByName) Swap(i, j int)
func (x MethodsByName) Less(i, j int) bool

// EmbeddedsByName sorts embedded types by name.
type EmbeddedsByName []*Field

func (x EmbeddedsByName) Len() int
func (x EmbeddedsByName) Swap(i, j int)
func (x EmbeddedsByName) Less(i, j int) bool
