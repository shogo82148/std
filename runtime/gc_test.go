// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

type Node1 struct {
	Value       [1]uintptr
	Left, Right *byte
}

type Node8 struct {
	Value       [8]uintptr
	Left, Right *byte
}

type Node64 struct {
	Value       [64]uintptr
	Left, Right *byte
}

type Node64Dead struct {
	Left, Right *byte
	Value       [64]uintptr
}

type Node124 struct {
	Value       [124]uintptr
	Left, Right *byte
}

type Node126 struct {
	Value       [126]uintptr
	Left, Right *byte
}

type Node128 struct {
	Value       [128]uintptr
	Left, Right *byte
}

type Node130 struct {
	Value       [130]uintptr
	Left, Right *byte
}

type Node1024 struct {
	Value       [1024]uintptr
	Left, Right *byte
}
