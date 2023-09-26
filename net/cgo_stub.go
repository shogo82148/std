// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !cgo || netgo
// +build !cgo netgo

// Stub cgo routines for systems that do not use cgo to do network lookups.

package net
