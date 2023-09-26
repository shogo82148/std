// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

var CallGC = &callGC

var GCBits = gcbits

type EmbedWithUnexpMeth struct{}

type OtherPkgFields struct {
	OtherExported   int
	otherUnexported int
}

type Buffer struct {
	buf []byte
}

var MethodValueCallCodePtr = methodValueCallCodePtr
