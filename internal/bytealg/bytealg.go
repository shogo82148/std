// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytealg

// MaxLen is the maximum length of the string to be searched for (argument b) in Index.
// If MaxLen is not 0, make sure MaxLen >= 4.
var MaxLen int

// IndexRabinKarp uses the Rabin-Karp search algorithm to return the index of the
// first occurrence of sep in s, or -1 if not present.
func IndexRabinKarp[T string | []byte](s, sep T) int

// LastIndexRabinKarp uses the Rabin-Karp search algorithm to return the last index of the
// occurrence of sep in s, or -1 if not present.
func LastIndexRabinKarp[T string | []byte](s, sep T) int

// MakeNoZero makes a slice of length n and capacity of at least n Bytes
// without zeroing the bytes (including the bytes between len and cap).
// It is the caller's responsibility to ensure uninitialized bytes
// do not leak to the end user.
func MakeNoZero(n int) []byte
