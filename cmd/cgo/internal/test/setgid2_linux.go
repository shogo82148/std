// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Stress test setgid and thread creation. A thread
// can get a SIGSETXID signal early on at thread
// initialization, causing crash. See issue 53374.

package cgotest
