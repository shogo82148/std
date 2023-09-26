// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// We used to use this code for Darwin, but according to issue #19314
// waitid returns if the process is stopped, even when using WEXITED.

//go:build linux

package os
