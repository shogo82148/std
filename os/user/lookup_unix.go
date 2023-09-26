// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (darwin || freebsd || linux) && cgo
// +build darwin freebsd linux
// +build cgo

package user

/*
#include <unistd.h>
#include <sys/types.h>
#include <pwd.h>
#include <stdlib.h>

static int mygetpwuid_r(int uid, struct passwd *pwd,
	char *buf, size_t buflen, struct passwd **result) {
 return getpwuid_r(uid, pwd, buf, buflen, result);
}
*/

// Current returns the current user.
func Current() (*User, error)

// Lookup looks up a user by username. If the user cannot be found,
// the returned error is of type UnknownUserError.
func Lookup(username string) (*User, error)

// LookupId looks up a user by userid. If the user cannot be found,
// the returned error is of type UnknownUserIdError.
func LookupId(uid string) (*User, error)
