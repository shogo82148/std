// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	"syscall"
	"testing"
)

type DLL struct {
	*syscall.DLL
	t *testing.T
}

// TODO(register args): Remove this once we switch to using the register
// calling convention by default, since this is redundant with the existing
// tests.
