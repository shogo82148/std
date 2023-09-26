// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package user allows user account lookups by name or id.

For most Unix systems, this package has two internal implementations of
resolving user and group ids to names. One is written in pure Go and
parses /etc/passwd and /etc/group. The other is cgo-based and relies on
the standard C library (libc) routines such as getpwuid_r and getgrnam_r.

When cgo is available, cgo-based (libc-backed) code is used by default.
This can be overriden by using osusergo build tag, which enforces
the pure Go implementation.
*/
package user

// User represents a user account.
type User struct {
	Uid string

	Gid string

	Username string

	Name string

	HomeDir string
}

// Group represents a grouping of users.
//
// On POSIX systems Gid contains a decimal number representing the group ID.
type Group struct {
	Gid  string
	Name string
}

// UnknownUserIdError is returned by LookupId when a user cannot be found.
type UnknownUserIdError int

func (e UnknownUserIdError) Error() string

// UnknownUserError is returned by Lookup when
// a user cannot be found.
type UnknownUserError string

func (e UnknownUserError) Error() string

// UnknownGroupIdError is returned by LookupGroupId when
// a group cannot be found.
type UnknownGroupIdError string

func (e UnknownGroupIdError) Error() string

// UnknownGroupError is returned by LookupGroup when
// a group cannot be found.
type UnknownGroupError string

func (e UnknownGroupError) Error() string
