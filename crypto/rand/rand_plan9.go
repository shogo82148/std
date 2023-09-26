// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Plan9 cryptographically secure pseudorandom number
// generator.

package rand

// reader is a new pseudorandom generator that seeds itself by
// reading from /dev/random. The Read method on the returned
// reader always returns the full amount asked for, or else it
// returns an error. The generator is a fast key erasure RNG.
