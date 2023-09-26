// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DNS packet assembly.  See RFC 1035.
//
// This is intended to support name resolution during Dial.
// It doesn't have to be blazing fast.
//
// Each message structure has a Walk method that is used by
// a generic pack/unpack routine. Thus, if in the future we need
// to define new message structs, no new pack/unpack/printing code
// needs to be written.
//
// The first half of this file defines the DNS message formats.
// The second half implements the conversion to and from wire format.
// A few of the structure elements have string tags to aid the
// generic pack/unpack routines.
//
// TODO(rsc):  There are enough names defined in this file that they're all
// prefixed with dns.  Perhaps put this in its own package later.

package net

// Wire constants.

// A dnsStruct describes how to iterate over its fields to emulate
// reflective marshalling.

// The wire format for the DNS packet header.

// DNS queries.

// DNS responses (resource records).
// There are many types of messages,
// but they all share the same header.

// Map of constructors for each RR wire type.

// A manually-unpacked version of (id, bits).
// This is in its own struct for easy printing.
