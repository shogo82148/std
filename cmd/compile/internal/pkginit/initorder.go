// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkginit

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// Static initialization phase.
// These values are stored in two bits in Node.flags.
const (
	InitNotStarted = iota
	InitDone
	InitPending
)

type InitOrder struct {
	// blocking maps initialization assignments to the assignments
	// that depend on it.
	blocking map[ir.Node][]ir.Node

	// ready is the queue of Pending initialization assignments
	// that are ready for initialization.
	ready declOrder

	order map[ir.Node]int
}
