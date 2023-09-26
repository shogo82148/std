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
