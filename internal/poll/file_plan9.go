// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

// FDMutex is an exported fdMutex, only for Plan 9.
type FDMutex struct {
	fdmu fdMutex
}

func (fdmu *FDMutex) Incref() bool

func (fdmu *FDMutex) Decref() bool

func (fdmu *FDMutex) IncrefAndClose() bool

func (fdmu *FDMutex) ReadLock() bool

func (fdmu *FDMutex) ReadUnlock() bool

func (fdmu *FDMutex) WriteLock() bool

func (fdmu *FDMutex) WriteUnlock() bool
