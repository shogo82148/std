// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/sync"
)

// A Context is an opaque type checking context. It may be used to share
// identical type instances across type-checked packages or calls to
// Instantiate. Contexts are safe for concurrent use.
//
// The use of a shared context does not guarantee that identical instances are
// deduplicated in all cases.
type Context struct {
	mu        sync.Mutex
	typeMap   map[string][]ctxtEntry
	nextID    int
	originIDs map[Type]int
}

// NewContext creates a new Context.
func NewContext() *Context
