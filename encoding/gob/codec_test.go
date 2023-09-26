// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob

// Guarantee encoding format by comparing some encodings to hand-written values
type EncodeT struct {
	x uint64
	b []byte
}

// The result of encoding a true boolean with field number 7

// The result of encoding a number 17 with field number 7

// The result of encoding a number 17+19i with field number 7

// The result of encoding "hello" with field number 7

// These three structures have the same data with different indirections
type T0 struct {
	A int
	B int
	C int
	D int
}
type T1 struct {
	A int
	B *int
	C **int
	D ***int
}
type T2 struct {
	A ***int
	B **int
	C *int
	D int
}

type RT0 struct {
	A int
	B string
	C float64
}
type RT1 struct {
	C      float64
	B      string
	A      int
	NotSet string
}

// Like an RT0 but with fields we'll ignore on the decode side.
type IT0 struct {
	A        int64
	B        string
	Ignore_d []int
	Ignore_e [3]float64
	Ignore_f bool
	Ignore_g string
	Ignore_h []byte
	Ignore_i *RT1
	Ignore_m map[string]int
	C        float64
}

type Indirect struct {
	A ***[3]int
	S ***[]int
	M ****map[string]int
}

type Direct struct {
	A [3]int
	S []int
	M map[string]int
}

// An interface with several implementations
type Squarer interface {
	Square() int
}

type Int int

type Float float64

type Vector []int

type Point struct {
	X, Y int
}

// A struct with interfaces in it.
type InterfaceItem struct {
	I             int
	Sq1, Sq2, Sq3 Squarer
	F             float64
	Sq            []Squarer
}

// The same struct without interfaces
type NoInterfaceItem struct {
	I int
	F float64
}

// A struct with all basic types, stored in interfaces.
type BasicInterfaceItem struct {
	Int, Int8, Int16, Int32, Int64      interface{}
	Uint, Uint8, Uint16, Uint32, Uint64 interface{}
	Float32, Float64                    interface{}
	Complex64, Complex128               interface{}
	Bool                                interface{}
	String                              interface{}
	Bytes                               interface{}
}

type String string

type PtrInterfaceItem struct {
	Str1 interface{}
	Str2 interface{}
}

type U struct {
	A int
	B string
	c float64
	D uint
}

// A type that won't be defined in the gob until we send it in an interface value.
type OnTheFly struct {
	A int
}

type DT struct {
	A     int
	B     string
	C     float64
	I     interface{}
	J     interface{}
	I_nil interface{}
	M     map[string]int
	T     [3]int
	S     []string
}
