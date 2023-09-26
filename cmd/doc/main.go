// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Doc (usually run as go doc) accepts zero, one or two arguments.
//
// Zero arguments:
//
//	go doc
//
// Show the documentation for the package in the current directory.
//
// One argument:
//
//	go doc <pkg>
//	go doc <sym>[.<method>]
//	go doc [<pkg>].<sym>[.<method>]
//
// The first item in this list that succeeds is the one whose documentation
// is printed. If there is a symbol but no package, the package in the current
// directory is chosen.
//
// Two arguments:
//
//	go doc <pkg> <sym>[.<method>]
//
// Show the documentation for the package, symbol, and method. The
// first argument must be a full package path. This is similar to the
// command-line usage for the godoc command.
//
// For commands, unless the -cmd flag is present "go doc command"
// shows only the package-level docs for the package.
//
// For complete documentation, run "go help doc".
package main
