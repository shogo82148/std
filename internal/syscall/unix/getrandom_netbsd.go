// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

// GetRandomFlag is a flag supported by the getrandom system call.
type GetRandomFlag uintptr

// GetRandom calls the getrandom system call.
func GetRandom(p []byte, flags GetRandomFlag) (n int, err error)
