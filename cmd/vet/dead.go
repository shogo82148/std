// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Simplified dead code detector. Used for skipping certain checks
// on unreachable code (for instance, shift checks on arch-specific code).
package main
