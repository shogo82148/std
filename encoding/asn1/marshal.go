// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asn1

// encoder represents a ASN.1 element that is waiting to be marshaled.

// Marshal returns the ASN.1 encoding of val.
//
// In addition to the struct tags recognised by Unmarshal, the following can be
// used:
//
//	ia5:		causes strings to be marshaled as ASN.1, IA5 strings
//	omitempty:	causes empty slices to be skipped
//	printable:	causes strings to be marshaled as ASN.1, PrintableString strings.
//	utf8:		causes strings to be marshaled as ASN.1, UTF8 strings
func Marshal(val interface{}) ([]byte, error)
