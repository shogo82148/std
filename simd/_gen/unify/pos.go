// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unify

type Pos struct {
	Path string
	Line int
}

func (p Pos) String() string

func (p Pos) AppendText(b []byte) ([]byte, error)
