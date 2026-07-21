// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssabase

// A Register is a machine register, like AX.
// They are numbered densely from 0 (for each architecture).
type Register struct {
	Num    int32
	ObjNum int16
	Name   string
}

func (r *Register) String() string
