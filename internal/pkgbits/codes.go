// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkgbits

// A Code is an enum value that can be encoded into bitstreams.
//
// Code types are preferable for enum types, because they allow
// Decoder to detect desyncs.
type Code interface {
	Marker() SyncMarker

	Value() int
}

// A CodeVal distinguishes among go/constant.Value encodings.
type CodeVal int

func (c CodeVal) Marker() SyncMarker
func (c CodeVal) Value() int

const (
	ValBool CodeVal = iota
	ValString
	ValInt64
	ValBigInt
	ValBigRat
	ValBigFloat
)

// A CodeType distinguishes among go/types.Type encodings.
type CodeType int

func (c CodeType) Marker() SyncMarker
func (c CodeType) Value() int

const (
	TypeBasic CodeType = iota
	TypeNamed
	TypePointer
	TypeSlice
	TypeArray
	TypeChan
	TypeMap
	TypeSignature
	TypeStruct
	TypeInterface
	TypeUnion
	TypeTypeParam
)

// A CodeObj distinguishes among go/types.Object encodings.
type CodeObj int

func (c CodeObj) Marker() SyncMarker
func (c CodeObj) Value() int

const (
	ObjAlias CodeObj = iota
	ObjConst
	ObjType
	ObjFunc
	ObjVar
	ObjStub
)
