// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js

package main

import (
	rtrace "runtime/trace"
)

// stacks is a fake stack map populated for test.
