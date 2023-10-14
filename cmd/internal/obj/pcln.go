// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package obj

// PCIter iterates over encoded pcdata tables.
type PCIter struct {
	p       []byte
	PC      uint32
	NextPC  uint32
	PCScale uint32
	Value   int32
	start   bool
	Done    bool
}

// NewPCIter creates a PCIter with a scale factor for the PC step size.
func NewPCIter(pcScale uint32) *PCIter

// Next advances it to the Next pc.
func (it *PCIter) Next()

// init prepares it to iterate over p,
// and advances it to the first pc.
func (it *PCIter) Init(p []byte)
