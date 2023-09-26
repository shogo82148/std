// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Windows cryptographically secure pseudorandom number
// generator.

package rand

// A rngReader satisfies reads by reading from the Windows CryptGenRandom API.
