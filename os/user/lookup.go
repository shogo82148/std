// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package user

// Current returns the current user.
//
// The first call will cache the current user information.
// Subsequent calls will return the cached value and will not reflect
// changes to the current user.
//
// On some systems, if the current user cannot be found in the user
// database, Current builds the returned user from the process's user and
// group IDs and from environment variables such as USER and HOME. In that
// case the returned user does not correspond to an entry in the user
// database.
func Current() (*User, error)

// Lookup looks up a user by username. If the user cannot be found, the
// returned error is of type [UnknownUserError].
//
// Lookup first calls [Current]. If that succeeds and the current user's
// name matches username, Lookup returns the current user without a separate
// lookup by name. The result may be one that [Current] built from the
// environment rather than from the user database.
func Lookup(username string) (*User, error)

// LookupId looks up a user by userid. If the user cannot be found, the
// returned error is of type [UnknownUserIdError].
//
// LookupId first calls [Current]. If that succeeds and the current user's
// ID matches uid, LookupId returns the current user without a separate
// lookup by ID. The result may be one that [Current] built from the
// environment rather than from the user database.
func LookupId(uid string) (*User, error)

// LookupGroup looks up a group by name. If the group cannot be found, the
// returned error is of type [UnknownGroupError].
func LookupGroup(name string) (*Group, error)

// LookupGroupId looks up a group by groupid. If the group cannot be found, the
// returned error is of type [UnknownGroupIdError].
func LookupGroupId(gid string) (*Group, error)

// GroupIds returns the list of group IDs that the user is a member of.
func (u *User) GroupIds() ([]string, error)
