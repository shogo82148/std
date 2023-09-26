// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob

// userTypeInfo stores the information associated with a type the user has handed
// to the package.  It's computed once and stored in a map keyed by reflection
// type.

// externalEncoding bits

// A typeId represents a gob Type as an integer that can be passed on the wire.
// Internally, typeIds are used as keys to a map to recover the underlying type info.

// CommonType holds elements of all types.
// It is a historical artifact, kept for binary compatibility and exported
// only for the benefit of the package's encoding of type descriptors. It is
// not intended for direct use by clients.
type CommonType struct {
	Name string
	Id   typeId
}

// Predefined because it's needed by the Decoder

// Array type

// GobEncoder type (something that implements the GobEncoder interface)

// Map type

// Slice type

// Struct type

// Representation of the information we send and receive about this type.
// Each value we send is preceded by its type definition: an encoded int.
// However, the very first time we send the value, we first send the pair
// (-id, wireType).
// For bootstrapping purposes, we assume that the recipient knows how
// to decode a wireType; it is exactly the wireType struct here, interpreted
// using the gob rules for sending a structure, except that we assume the
// ids for wireType and structType etc. are known.  The relevant pieces
// are built in encode.go's init() function.
// To maintain binary compatibility, if you extend this type, always put
// the new fields last.

// typeInfoMap is an atomic pointer to map[reflect.Type]*typeInfo.
// It's updated copy-on-write. Readers just do an atomic load
// to get the current version of the map. Writers make a full copy of
// the map and atomically update the pointer to point to the new map.
// Under heavy read contention, this is significantly faster than a map
// protected by a mutex.

// GobEncoder is the interface describing data that provides its own
// representation for encoding values for transmission to a GobDecoder.
// A type that implements GobEncoder and GobDecoder has complete
// control over the representation of its data and may therefore
// contain things such as private fields, channels, and functions,
// which are not usually transmissible in gob streams.
//
// Note: Since gobs can be stored permanently, It is good design
// to guarantee the encoding used by a GobEncoder is stable as the
// software evolves.  For instance, it might make sense for GobEncode
// to include a version number in the encoding.
type GobEncoder interface {
	GobEncode() ([]byte, error)
}

// GobDecoder is the interface describing data that provides its own
// routine for decoding transmitted values sent by a GobEncoder.
type GobDecoder interface {
	GobDecode([]byte) error
}

// RegisterName is like Register but uses the provided name rather than the
// type's default.
func RegisterName(name string, value interface{})

// Register records a type, identified by a value for that type, under its
// internal type name.  That name will identify the concrete type of a value
// sent or received as an interface variable.  Only types that will be
// transferred as implementations of interface values need to be registered.
// Expecting to be used only during initialization, it panics if the mapping
// between types and names is not a bijection.
func Register(value interface{})
