// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maps

// runtime variable to check if the processor we're running on
// actually supports the instructions used by the AES-based
// hash implementation.
var UseAeshash bool

func AlgInit()
