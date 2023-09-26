// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Parse "zoneinfo" time zone file.
// This is a fairly standard file format used on OS X, Linux, BSD, Sun, and others.
// See tzfile(5), https://en.wikipedia.org/wiki/Zoneinfo,
// and ftp://munnari.oz.au/pub/oldtz/

package time

// loadFromEmbeddedTZData is used to load a specific tzdata file
// from tzdata information embedded in the binary itself.
// This is set when the time/tzdata package is imported,
// via registerLoadFromEmbeddedTzdata.

// maxFileSize is the max permitted size of files read by readFile.
// As reference, the zoneinfo.zip distributed by Go is ~350 KB,
// so 10MB is overkill.

// Copies of io.Seek* constants to avoid importing "io":

// Simple I/O interface to binary blob of data.

// LoadLocationFromTZData returns a Location with the given name
// initialized from the IANA Time Zone database-formatted data.
// The data should be in the format of a standard IANA time zone file
// (for example, the content of /etc/localtime on Unix systems).
func LoadLocationFromTZData(name string, data []byte) (*Location, error)

// loadTzinfoFromTzdata returns the time zone information of the time zone
// with the given name, from a tzdata database file as they are typically
// found on android.
