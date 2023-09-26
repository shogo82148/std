// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.arenas

package arena_test

type T1 struct {
	n int
}
type T2 [1 << 20]byte
