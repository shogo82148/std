// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// Valid DNS SRV reply

// Corrupt DNS SRV reply, with its final RR having a bogus length
// (perhaps it was truncated, or it's malicious) The mutation is the
// capital "FF" below, instead of the proper "21".

// TXT reply with one <character-string>

// TXT reply with more than one <character-string>.
// See https://tools.ietf.org/html/rfc1035#section-3.3.14

// DataLength field should be sum of all TXT fields. In this case it's less.

// Same as above but DataLength is more than sum of TXT fields.

// TXT Length field is less than actual length.

// TXT Length field is more than actual length.
