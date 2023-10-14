// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lockedfile

import (
	"github.com/shogo82148/std/sync"
)

// A Mutex provides mutual exclusion within and across processes by locking a
// well-known file. Such a file generally guards some other part of the
// filesystem: for example, a Mutex file in a directory might guard access to
// the entire tree rooted in that directory.
//
// Mutex does not implement sync.Locker: unlike a sync.Mutex, a lockedfile.Mutex
// can fail to lock (e.g. if there is a permission error in the filesystem).
//
// Like a sync.Mutex, a Mutex may be included as a field of a larger struct but
// must not be copied after first use. The Path field must be set before first
// use and must not be change thereafter.
type Mutex struct {
	Path string
	mu   sync.Mutex
}

// MutexAt returns a new Mutex with Path set to the given non-empty path.
func MutexAt(path string) *Mutex

func (mu *Mutex) String() string

// Lock attempts to lock the Mutex.
//
// If successful, Lock returns a non-nil unlock function: it is provided as a
// return-value instead of a separate method to remind the caller to check the
// accompanying error. (See https://golang.org/issue/20803.)
func (mu *Mutex) Lock() (unlock func(), err error)
