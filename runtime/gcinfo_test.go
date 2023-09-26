// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

const (
	BitsDead = iota
	BitsScalar
	BitsPointer
	BitsMultiWord
)

const (
	BitsString = iota
	BitsSlice
	BitsIface
	BitsEface
)

type ScalarPtr struct {
	q int
	w *int
	e int
	r *int
	t int
	y *int
}

type PtrScalar struct {
	q *int
	w int
	e *int
	r int
	t *int
	y int
}

type BigStruct struct {
	q *int
	w byte
	e [17]byte
	r []byte
	t int
	y uint16
	u uint64
	i string
}

type Iface interface {
	f()
}

type IfaceImpl int
