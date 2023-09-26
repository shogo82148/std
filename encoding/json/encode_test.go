// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

type Optionals struct {
	Sr string `json:"sr"`
	So string `json:"so,omitempty"`
	Sw string `json:"-"`

	Ir int `json:"omitempty"`
	Io int `json:"io,omitempty"`

	Slr []string `json:"slr,random"`
	Slo []string `json:"slo,omitempty"`

	Mr map[string]interface{} `json:"mr"`
	Mo map[string]interface{} `json:",omitempty"`

	Fr float64 `json:"fr"`
	Fo float64 `json:"fo,omitempty"`

	Br bool `json:"br"`
	Bo bool `json:"bo,omitempty"`

	Ur uint `json:"ur"`
	Uo uint `json:"uo,omitempty"`

	Str struct{} `json:"str"`
	Sto struct{} `json:"sto,omitempty"`
}

type StringTag struct {
	BoolStr bool   `json:",string"`
	IntStr  int64  `json:",string"`
	StrStr  string `json:",string"`
}

// byte slices are special even if they're renamed types.

// Ref has Marshaler and Unmarshaler methods with pointer receiver.
type Ref int

// Val has Marshaler methods with value receiver.
type Val int

// RefText has Marshaler and Unmarshaler methods with pointer receiver.
type RefText int

// ValText has Marshaler methods with value receiver.
type ValText int

// C implements Marshaler and returns unescaped JSON.
type C int

// CText implements Marshaler and returns unescaped text.
type CText int

type BugA struct {
	S string
}

type BugB struct {
	BugA
	S string
}

type BugC struct {
	S string
}

// Legal Go: We never use the repeated embedded field (S).
type BugX struct {
	A int
	BugA
	BugB
}

// Issue 16042. Even if a nil interface value is passed in
// as long as it implements MarshalJSON, it should be marshaled.

type BugD struct {
	XXX string `json:"S"`
}

// BugD's tagged S field should dominate BugA's.
type BugY struct {
	BugA
	BugD
}

// There are no tags here, so S should not appear.
type BugZ struct {
	BugA
	BugC
	BugY
}

// syntactic checks on form of marshaled floating point numbers.
