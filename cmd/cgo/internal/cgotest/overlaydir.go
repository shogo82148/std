// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgotest

// OverlayDir makes a minimal-overhead copy of srcRoot in which new files may be added.
func OverlayDir(dstRoot, srcRoot string) error
