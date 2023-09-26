// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// res_nsearch, for cgo systems where that's available.

//go:build cgo && !netgo && unix && !(darwin || linux || openbsd)

package net

/*
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <unistd.h>
#include <string.h>
#include <arpa/nameser.h>
#include <resolv.h>

#cgo !aix,!dragonfly,!freebsd LDFLAGS: -lresolv
*/
