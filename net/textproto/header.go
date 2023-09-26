// Copyright 2010 The Go Authors.  All rights reserved.
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
// the single element value.  It replaces any existing
// values associated with key.
func (h MIMEHeader) Set(key, value string)

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns "".
// Get is a convenience method.  For more complex queries,
// access the map directly.
func (h MIMEHeader) Get(key string) string

// Del deletes the values associated with key.
func (h MIMEHeader) Del(key string)
