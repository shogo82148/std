// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !cgo && !windows
// +build !cgo,!windows

package user

func Current() (*User, error)

func Lookup(username string) (*User, error)

func LookupId(string) (*User, error)
