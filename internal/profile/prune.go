// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Implements methods to remove frames from profiles.

package profile

import (
	"github.com/shogo82148/std/regexp"
)

// Prune removes all nodes beneath a node matching dropRx, and not
// matching keepRx. If the root node of a Sample matches, the sample
// will have an empty stack.
func (p *Profile) Prune(dropRx, keepRx *regexp.Regexp)

// RemoveUninteresting prunes and elides profiles using built-in
// tables of uninteresting function names.
func (p *Profile) RemoveUninteresting() error
