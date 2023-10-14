// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

// GetRandomFlag is a flag supported by the getrandom system call.
type GetRandomFlag uintptr

const (
	// GRND_NONBLOCK means return EAGAIN rather than blocking.
	GRND_NONBLOCK GetRandomFlag = 0x0001

	// GRND_RANDOM means use the /dev/random pool instead of /dev/urandom.
	GRND_RANDOM GetRandomFlag = 0x0002
)

// GetRandom calls the getrandom system call.
func GetRandom(p []byte, flags GetRandomFlag) (n int, err error)
