// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

type InternalExample struct {
	Name      string
	F         func()
	Output    string
	Unordered bool
}

// RunExamples is an internal function but exported because it is cross-package;
// it is part of the implementation of the "go test" command.
func RunExamples(matchString func(pat, str string) (bool, error), examples []InternalExample) (ok bool)
