// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// A rwmutex is a reader/writer mutual exclusion lock.
// The lock can be held by an arbitrary number of readers or a single writer.
// This is a variant of sync.RWMutex, for the runtime package.
// Like mutex, rwmutex blocks the calling M.
// It does not interact with the goroutine scheduler.
