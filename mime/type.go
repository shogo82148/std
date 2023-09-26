// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mime implements parts of the MIME spec.
package mime

// TypeByExtension returns the MIME type associated with the file extension ext.
// The extension ext should begin with a leading dot, as in ".html".
// When ext has no associated type, TypeByExtension returns "".
//
// The built-in table is small but on unix it is augmented by the local
// system's mime.types file(s) if available under one or more of these
// names:
//
//	/etc/mime.types
//	/etc/apache2/mime.types
//	/etc/apache/mime.types
//
// Windows system mime types are extracted from registry.
//
// Text types have the charset parameter set to "utf-8" by default.
func TypeByExtension(ext string) string

// AddExtensionType sets the MIME type associated with
// the extension ext to typ.  The extension should begin with
// a leading dot, as in ".html".
func AddExtensionType(ext, typ string) error
