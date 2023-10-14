// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"github.com/shogo82148/std/cmd/link/internal/loader"
	"github.com/shogo82148/std/sync"
)

// ErrorReporter is used to make error reporting thread safe.
type ErrorReporter struct {
	loader.ErrorReporter
	unresSyms  map[unresolvedSymKey]bool
	unresMutex sync.Mutex
	SymName    symNameFn
}
