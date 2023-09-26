// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !netgo && (darwin || dragonfly || freebsd || linux || netbsd || openbsd)
// +build !netgo
// +build darwin dragonfly freebsd linux netbsd openbsd

package net

/*
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
*/
