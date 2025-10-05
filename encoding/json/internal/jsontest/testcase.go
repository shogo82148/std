// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontest

// CaseName is a case name annotated with a file and line.
type CaseName struct {
	Name  string
	Where CasePos
}

// Name annotates a case name with the file and line of the caller.
func Name(s string) (c CaseName)

// CasePos represents a file and line number.
type CasePos struct{ pc [1]uintptr }

func (pos CasePos) String() string
