// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

func AtExit(f func())

func Exit(code int)

// To enable tracing support (-t flag), set EnableTrace to true.
const EnableTrace = false
