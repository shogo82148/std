// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains extra hooks for testing the go command.
// It is compiled into the Go binary only when building the
// test copy; it does not get compiled into the standard go
// command, so these testing hooks are not present in the
// go command that everyone uses.

//go:build testgo
// +build testgo

package main
