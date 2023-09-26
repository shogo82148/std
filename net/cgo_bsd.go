// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !netgo && (darwin || dragonfly || freebsd)
// +build !netgo
// +build darwin dragonfly freebsd

package net

/*
#include <netdb.h>
*/
