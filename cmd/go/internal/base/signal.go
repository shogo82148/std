// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

// Interrupted is closed when the go command receives an interrupt signal.
var Interrupted = make(chan struct{})

// StartSigHandlers starts the signal handlers.
func StartSigHandlers()
