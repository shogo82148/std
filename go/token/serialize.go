// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

// Read calls decode to deserialize a file set into s; s must not be nil.
func (s *FileSet) Read(decode func(interface{}) error) error

// Write calls encode to serialize the file set s.
func (s *FileSet) Write(encode func(interface{}) error) error
