// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textproto

// A MIMEHeader represents a MIME-style header mapping
// keys to sets of values.
type MIMEHeader map[string][]string

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
func (h MIMEHeader) Add(key, value string)

// Set sets the header entries associated with key to
// the single element value. It replaces any existing
// values associated with key.
func (h MIMEHeader) Set(key, value string)

// Get gets the first value associated with the given key.
// It is case insensitive; [CanonicalMIMEHeaderKey] is used
// to canonicalize the provided key.
// If there are no values associated with the key, Get returns "".
// To use non-canonical keys, access the map directly.
func (h MIMEHeader) Get(key string) string

// Values returns all values associated with the given key.
// It is case insensitive; [CanonicalMIMEHeaderKey] is
// used to canonicalize the provided key. To use non-canonical
// keys, access the map directly.
// The returned slice is not a copy.
func (h MIMEHeader) Values(key string) []string

// Del deletes the values associated with key.
func (h MIMEHeader) Del(key string)
