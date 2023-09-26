// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo && !netgo && (android || linux || solaris)
// +build cgo
// +build !netgo
// +build android linux solaris

package net

/*
#include <sys/types.h>
#include <sys/socket.h>

#include <netinet/in.h>
*/
