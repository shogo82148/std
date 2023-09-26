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

// C implements Marshaler and returns unescaped JSON.
type C int
