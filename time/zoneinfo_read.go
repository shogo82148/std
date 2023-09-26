// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Parse "zoneinfo" time zone file.
// This is a fairly standard file format used on OS X, Linux, BSD, Sun, and others.
// See tzfile(5), http://en.wikipedia.org/wiki/Zoneinfo,
// and ftp://munnari.oz.au/pub/oldtz/

package time

// maxFileSize is the max permitted size of files read by readFile.
// As reference, the zoneinfo.zip distributed by Go is ~350 KB,
// so 10MB is overkill.

// Copies of io.Seek* constants to avoid importing "io":

// Simple I/O interface to binary blob of data.
