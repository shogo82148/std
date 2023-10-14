// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buildid

// ReadFile reads the build ID from an archive or executable file.
func ReadFile(name string) (id string, err error)

// HashToString converts the hash h to a string to be recorded
// in package archives and binaries as part of the build ID.
// We use the first 120 bits of the hash (5 chunks of 24 bits each) and encode
// it in base64, resulting in a 20-byte string. Because this is only used for
// detecting the need to rebuild installed files (not for lookups
// in the object file cache), 120 bits are sufficient to drive the
// probability of a false "do not need to rebuild" decision to effectively zero.
// We embed two different hashes in archives and four in binaries,
// so cutting to 20 bytes is a significant savings when build IDs are displayed.
// (20*4+3 = 83 bytes compared to 64*4+3 = 259 bytes for the
// more straightforward option of printing the entire h in base64).
func HashToString(h [32]byte) string
