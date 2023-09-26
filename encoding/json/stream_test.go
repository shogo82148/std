// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

// CaseName is a case name annotated with a file and line.
type CaseName struct {
	Name  string
	Where CasePos
}

// CasePos represents a file and line number.
type CasePos struct{ pc [1]uintptr }

// Test values for the stream test.
// One of each JSON kind.
