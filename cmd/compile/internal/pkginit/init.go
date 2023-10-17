// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkginit

// MakeTask makes an initialization record for the package, if necessary.
// See runtime/proc.go:initTask for its layout.
// The 3 tasks for initialization are:
//  1. Initialize all of the packages the current package depends on.
//  2. Initialize all the variables that have initializers.
//  3. Run any init functions.
func MakeTask()
