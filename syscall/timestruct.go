// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

package syscall

// TimespecToNsecは、tsに格納されている時間をナノ秒で返します。
func TimespecToNsec(ts Timespec) int64

// NsecToTimespecは、ナノ秒数を[Timespec]に変換します。
func NsecToTimespec(nsec int64) Timespec

// TimevalToNsecは、tvに格納されている時間をナノ秒で返します。
func TimevalToNsec(tv Timeval) int64

// NsecToTimevalは、ナノ秒数を[Timeval]に変換します。
func NsecToTimeval(nsec int64) Timeval
