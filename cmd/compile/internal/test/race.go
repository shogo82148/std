// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !compiler_bootstrap

package test

// Issue 55357: data race when building multiple instantiations of
// generic closures with _ parameters.
func Issue55357()

type T55357[T any] struct{}

func (q *T55357[T]) Count() (n int, rerr error)

func (q *T55357[T]) List() (list []T, rerr error)
