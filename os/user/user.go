// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package user allows user account lookups by name or id.

For most Unix systems, this package has two internal implementations of
resolving user and group ids to names, and listing supplementary group IDs.
One is written in pure Go and parses /etc/passwd and /etc/group. The other
is cgo-based and relies on the standard C library (libc) routines such as
getpwuid_r, getgrnam_r, and getgrouplist.

When cgo is available, and the required routines are implemented in libc
for a particular platform, cgo-based (libc-backed) code is used.
This can be overridden by using osusergo build tag, which enforces
the pure Go implementation.
*/
package user

// User represents a user account.
type User struct {
	// Uid is the user ID.
	// On POSIX systems, this is a decimal number representing the uid.
	// On Windows, this is a security identifier (SID) in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Uid string
	// Gid is the primary group ID.
	// On POSIX systems, this is a decimal number representing the gid.
	// On Windows, this is a SID in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Gid string
	// Username is the login name.
	Username string
	// Name is the user's real or display name.
	// It might be blank.
	// On POSIX systems, this is the first (or only) entry in the GECOS field
	// list.
	// On Windows, this is the user's display name.
	// On Plan 9, this is the contents of /dev/user.
	Name string
	// HomeDir is the path to the user's home directory (if they have one).
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
