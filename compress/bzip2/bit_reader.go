// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bzip2

// bitReader wraps an io.Reader and provides the ability to read values,
// bit-by-bit, from it. Its Read* methods don't return the usual error
// because the error handling was verbose. Instead, any error is kept and can
// be checked afterwards.

// bitReader needs to read bytes from an io.Reader. We attempt to convert the
// given io.Reader to this interface and, if it doesn't already fit, we wrap in
// a bufio.Reader.
