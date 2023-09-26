// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package gob manages streams of gobs - binary values exchanged between an
Encoder (transmitter) and a Decoder (receiver).  A typical use is transporting
arguments and results of remote procedure calls (RPCs) such as those provided by
package "rpc".

A stream of gobs is self-describing.  Each data item in the stream is preceded by
a specification of its type, expressed in terms of a small set of predefined
types.  Pointers are not transmitted, but the things they point to are
transmitted; that is, the values are flattened.  Recursive types work fine, but
recursive values (data with cycles) are problematic.  This may change.

To use gobs, create an Encoder and present it with a series of data items as
values or addresses that can be dereferenced to values.  The Encoder makes sure
all type information is sent before it is needed.  At the receive side, a
Decoder retrieves values from the encoded stream and unpacks them into local
variables.

The source and destination values/types need not correspond exactly.  For structs,
fields (identified by name) that are in the source but absent from the receiving
variable will be ignored.  Fields that are in the receiving variable but missing
from the transmitted type or value will be ignored in the destination.  If a field
with the same name is present in both, their types must be compatible. Both the
receiver and transmitter will do all necessary indirection and dereferencing to
convert between gobs and actual Go values.  For instance, a gob type that is
schematically,

	struct { A, B int }

can be sent from or received into any of these Go types:

	struct { A, B int }	// the same
	*struct { A, B int }	// extra indirection of the struct
	struct { *A, **B int }	// extra indirection of the fields
	struct { A, B int64 }	// different concrete value type; see below

It may also be received into any of these:

	struct { A, B int }	// the same
	struct { B, A int }	// ordering doesn't matter; matching is by name
	struct { A, B, C int }	// extra field (C) ignored
	struct { B int }	// missing field (A) ignored; data will be dropped
	struct { B, C int }	// missing field (A) ignored; extra field (C) ignored.

Attempting to receive into these types will draw a decode error:

	struct { A int; B uint }	// change of signedness for B
	struct { A int; B float }	// change of type for B
	struct { }			// no field names in common
	struct { C, D int }		// no field names in common

Integers are transmitted two ways: arbitrary precision signed integers or
arbitrary precision unsigned integers.  There is no int8, int16 etc.
discrimination in the gob format; there are only signed and unsigned integers.  As
described below, the transmitter sends the value in a variable-length encoding;
the receiver accepts the value and stores it in the destination variable.
Floating-point numbers are always sent using IEEE-754 64-bit precision (see
below).

Signed integers may be received into any signed integer variable: int, int16, etc.;
unsigned integers may be received into any unsigned integer variable; and floating
point values may be received into any floating point variable.  However,
the destination variable must be able to represent the value or the decode
operation will fail.

Structs, arrays and slices are also supported.  Strings and arrays of bytes are
supported with a special, efficient representation (see below).  When a slice is
decoded, if the existing slice has capacity the slice will be extended in place;
if not, a new array is allocated.  Regardless, the length of the resulting slice
reports the number of elements decoded.

Functions and channels cannot be sent in a gob.  Attempting
to encode a value that contains one will fail.

The rest of this comment documents the encoding, details that are not important
for most users.  Details are presented bottom-up.

An unsigned integer is sent one of two ways.  If it is less than 128, it is sent
as a byte with that value.  Otherwise it is sent as a minimal-length big-endian
(high byte first) byte stream holding the value, preceded by one byte holding the
byte count, negated.  Thus 0 is transmitted as (00), 7 is transmitted as (07) and
256 is transmitted as (FE 01 00).

A boolean is encoded within an unsigned integer: 0 for false, 1 for true.

A signed integer, i, is encoded within an unsigned integer, u.  Within u, bits 1
upward contain the value; bit 0 says whether they should be complemented upon
receipt.  The encode algorithm looks like this:

	uint u;
	if i < 0 {
		u = (^i << 1) | 1	// complement i, bit 0 is 1
	} else {
		u = (i << 1)	// do not complement i, bit 0 is 0
	}
	encodeUnsigned(u)

The low bit is therefore analogous to a sign bit, but making it the complement bit
instead guarantees that the largest negative integer is not a special case.  For
example, -129=^128=(^256>>1) encodes as (FE 01 01).

Floating-point numbers are always sent as a representation of a float64 value.
That value is converted to a uint64 using math.Float64bits.  The uint64 is then
byte-reversed and sent as a regular unsigned integer.  The byte-reversal means the
exponent and high-precision part of the mantissa go first.  Since the low bits are
often zero, this can save encoding bytes.  For instance, 17.0 is encoded in only
three bytes (FE 31 40).

Strings and slices of bytes are sent as an unsigned count followed by that many
uninterpreted bytes of the value.

All other slices and arrays are sent as an unsigned count followed by that many
elements using the standard gob encoding for their type, recursively.

Maps are sent as an unsigned count followed by that many key, element
pairs. Empty but non-nil maps are sent, so if the sender has allocated
a map, the receiver will allocate a map even if no elements are
transmitted.

Structs are sent as a sequence of (field number, field value) pairs.  The field
value is sent using the standard gob encoding for its type, recursively.  If a
field has the zero value for its type, it is omitted from the transmission.  The
field number is defined by the type of the encoded struct: the first field of the
encoded type is field 0, the second is field 1, etc.  When encoding a value, the
field numbers are delta encoded for efficiency and the fields are always sent in
order of increasing field number; the deltas are therefore unsigned.  The
initialization for the delta encoding sets the field number to -1, so an unsigned
integer field 0 with value 7 is transmitted as unsigned delta = 1, unsigned value
= 7 or (01 07).  Finally, after all the fields have been sent a terminating mark
denotes the end of the struct.  That mark is a delta=0 value, which has
representation (00).

Interface types are not checked for compatibility; all interface types are
treated, for transmission, as members of a single "interface" type, analogous to
int or []byte - in effect they're all treated as interface{}.  Interface values
are transmitted as a string identifying the concrete type being sent (a name
that must be pre-defined by calling Register), followed by a byte count of the
length of the following data (so the value can be skipped if it cannot be
stored), followed by the usual encoding of concrete (dynamic) value stored in
the interface value.  (A nil interface value is identified by the empty string
and transmits no value.) Upon receipt, the decoder verifies that the unpacked
concrete item satisfies the interface of the receiving variable.

The representation of types is described below.  When a type is defined on a given
connection between an Encoder and Decoder, it is assigned a signed integer type
id.  When Encoder.Encode(v) is called, it makes sure there is an id assigned for
the type of v and all its elements and then it sends the pair (typeid, encoded-v)
where typeid is the type id of the encoded type of v and encoded-v is the gob
encoding of the value v.

To define a type, the encoder chooses an unused, positive type id and sends the
pair (-type id, encoded-type) where encoded-type is the gob encoding of a wireType
description, constructed from these types:

	type wireType struct {
		ArrayT  *ArrayType
		SliceT  *SliceType
		StructT *StructType
		MapT    *MapType
	}
	type arrayType struct {
		CommonType
		Elem typeId
		Len  int
	}
	type CommonType struct {
		Name string // the name of the struct type
		Id  int    // the id of the type, repeated so it's inside the type
	}
	type sliceType struct {
		CommonType
		Elem typeId
	}
	type structType struct {
		CommonType
		Field []*fieldType // the fields of the struct.
	}
	type fieldType struct {
		Name string // the name of the field.
		Id   int    // the type id of the field, which must be already defined
	}
	type mapType struct {
		CommonType
		Key  typeId
		Elem typeId
	}

If there are nested type ids, the types for all inner type ids must be defined
before the top-level type id is used to describe an encoded-v.

For simplicity in setup, the connection is defined to understand these types a
priori, as well as the basic gob types int, uint, etc.  Their ids are:

	bool        1
	int         2
	uint        3
	float       4
	[]byte      5
	string      6
	complex     7
	interface   8
	// gap for reserved ids.
	WireType    16
	ArrayType   17
	CommonType  18
	SliceType   19
	StructType  20
	FieldType   21
	// 22 is slice of fieldType.
	MapType     23

Finally, each message created by a call to Encode is preceded by an encoded
unsigned integer count of the number of bytes remaining in the message.  After
the initial type name, interface values are wrapped the same way; in effect, the
interface value acts like a recursive invocation of Encode.

In summary, a gob stream looks like

	(byteCount (-type id, encoding of a wireType)* (type id, encoding of a value))*

where * signifies zero or more repetitions and the type id of a value must
be predefined or be defined before the value in the stream.

See "Gobs of data" for a design discussion of the gob wire format:
http://golang.org/doc/articles/gobs_of_data.html
*/
package gob
