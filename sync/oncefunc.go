// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

// OnceFuncは関数fを一度だけ呼び出す関数を返します。返された関数は並行して呼び出すことができます。
//
// fがパニックを起こした場合、返された関数はすべての呼び出しで同じ値でパニックを起こします。
func OnceFunc(f func()) func()

// OnceValue は、関数 f を一度だけ呼び出し、f の戻り値を返す関数を返します。
// 返された関数は、同時に呼び出すことができます。
//
// f がパニックを起こした場合、返された関数はすべての呼び出しで同じ値を持つパニックを発生させます。
func OnceValue[T any](f func() T) func() T

// OnceValuesは、fを一度だけ呼び出し、fによって返された値を返す関数を返します。返された関数は並行して呼び出すことができます。
//
// fがパニックを引き起こした場合、返された関数はすべての呼び出しで同じ値でパニックを起こします。
func OnceValues[T1, T2 any](f func() (T1, T2)) func() (T1, T2)
