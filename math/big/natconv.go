// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements nat-to-string conversion functions.

package big

// MaxBase is the largest number base accepted for string conversions.
const MaxBase = 'z' - 'a' + 10 + 1

// Split blocks greater than leafSize Words (or set to 0 to disable recursive conversion)
// Benchmark and configure leafSize using: go test -bench="Leaf"
//   8 and 16 effective on 3.0 GHz Xeon "Clovertown" CPU (128 byte cache lines)
//   8 and 16 effective on 2.66 GHz Core 2 Duo "Penryn" CPU
