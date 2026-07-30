// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The export command writes type information of Go programs.
// It should not be relied upon except by go list -export.
//
// It supports the following command-line protocol:
//
//	-V=full			print tool version
//	unit.cfg		description of unit to export
package main
