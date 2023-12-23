// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements nat-to-string conversion functions.

package big

// MaxBaseは、文字列変換に受け入れられる最大の数値基数です。
const MaxBase = 10 + ('z' - 'a' + 1) + ('Z' - 'A' + 1)
