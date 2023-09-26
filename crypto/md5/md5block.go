// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// MD5 block step.
// In its own file so that a faster assembly or C version
// can be substituted easily.

package md5

// table[i] = int((1<<32) * abs(sin(i+1 radians))).
