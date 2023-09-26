// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && cgo
// +build linux,cgo

// On systems that use glibc, calling malloc can create a new arena,
// and creating a new arena can read /sys/devices/system/cpu/online.
// If we are using cgo, we will call malloc when creating a new thread.
// That can break TestExtraFiles if we create a new thread that creates
// a new arena and opens the /sys file while we are checking for open
// file descriptors. Work around the problem by creating threads up front.
// See issue 25628.

package exec_test
