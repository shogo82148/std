// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/reflect"
)

const (
	// A generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

// Marshal returns the XML encoding of v.
//
// Marshal handles an array or slice by marshalling each of the elements.
// Marshal handles a pointer by marshalling the value it points at or, if the
// pointer is nil, by writing nothing.  Marshal handles an interface value by
// marshalling the value it contains or, if the interface value is nil, by
// writing nothing.  Marshal handles all other data by writing one or more XML
// elements containing the data.
//
// The name for the XML elements is taken from, in order of preference:
//   - the tag on the XMLName field, if the data is a struct
//   - the value of the XMLName field of type xml.Name
//   - the tag of the struct field used to obtain the data
//   - the name of the struct field used to obtain the data
//   - the name of the marshalled type
//
// The XML element for a struct contains marshalled elements for each of the
// exported fields of the struct, with these exceptions:
//   - the XMLName field, described above, is omitted.
//   - a field with tag "-" is omitted.
//   - a field with tag "name,attr" becomes an attribute with
//     the given name in the XML element.
//   - a field with tag ",attr" becomes an attribute with the
//     field name in the XML element.
//   - a field with tag ",chardata" is written as character data,
//     not as an XML element.
//   - a field with tag ",innerxml" is written verbatim, not subject
//     to the usual marshalling procedure.
//   - a field with tag ",comment" is written as an XML comment, not
//     subject to the usual marshalling procedure. It must not contain
//     the "--" string within it.
//   - a field with a tag including the "omitempty" option is omitted
//     if the field value is empty. The empty values are false, 0, any
//     nil pointer or interface value, and any array, slice, map, or
//     string of length zero.
//   - an anonymous struct field is handled as if the fields of its
//     value were part of the outer struct.
//
// If a field uses a tag "a>b>c", then the element c will be nested inside
// parent elements a and b.  Fields that appear next to each other that name
// the same parent will be enclosed in one XML element.
//
// See MarshalIndent for an example.
//
// Marshal will return an error if asked to marshal a channel, function, or map.
func Marshal(v interface{}) ([]byte, error)

// MarshalIndent works like Marshal, but each XML element begins on a new
// indented line that starts with prefix and is followed by one or more
// copies of indent according to the nesting depth.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

// An Encoder writes XML data to an output stream.
type Encoder struct {
	printer
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder

// Indent sets the encoder to generate XML in which each element
// begins on a new indented line that starts with prefix and is followed by
// one or more copies of indent according to the nesting depth.
func (enc *Encoder) Indent(prefix, indent string)

// Encode writes the XML encoding of v to the stream.
//
// See the documentation for Marshal for details about the conversion
// of Go values to XML.
func (enc *Encoder) Encode(v interface{}) error

// A MarshalXMLError is returned when Marshal encounters a type
// that cannot be converted into XML.
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string
