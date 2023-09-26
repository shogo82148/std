// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// Valid DNS SRV reply

// Corrupt DNS SRV reply, with its final RR having a bogus length
// (perhaps it was truncated, or it's malicious) The mutation is the
// capital "FF" below, instead of the proper "21".
