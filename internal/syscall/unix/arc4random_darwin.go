// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

// ARC4Random calls the macOS arc4random_buf(3) function.
func ARC4Random(p []byte)
