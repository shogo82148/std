// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package user

// Current returns the current user.
func Current() (*User, error)

// Lookup looks up a user by username.
func Lookup(username string) (*User, error)

// LookupId looks up a user by userid.
func LookupId(uid string) (*User, error)
