// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Fake networking for js/wasm and wasip1/wasm.
// It is intended to allow tests of other package to pass.

//go:build js || wasip1

package net

// A packetQueue is a set of 1-buffered channels implementing a FIFO queue
// of packets.
