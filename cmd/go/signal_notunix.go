// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build plan9 || windows
// +build plan9 windows

package main

// signalTrace is the signal to send to make a Go program
// crash with a stack trace.
