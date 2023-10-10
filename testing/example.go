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

// RunExamples は内部関数ですが、クロスパッケージとなっているため公開されています。
// これは「go test」コマンドの実装の一部です。
func RunExamples(matchString func(pat, str string) (bool, error), examples []InternalExample) (ok bool)
