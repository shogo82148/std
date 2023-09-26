// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !nacl && !js && !plan9 && !windows
// +build !nacl,!js,!plan9,!windows

package os_test

// noDeadline is a zero time.Time value, which cancels a deadline.
