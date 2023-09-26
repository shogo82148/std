// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test broken pipes on Unix systems.
//go:build !plan9 && !nacl && !js
// +build !plan9,!nacl,!js

package os_test

import (
	osexec "os/exec"
)
