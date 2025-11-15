// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains stub functions that are not meant to be called directly,
// but that will be assembled together using the inlining logic in runtime/_mkmalloc
// to produce a full mallocgc function that's specialized for a span class
// or specific size in the case of the tiny allocator.
//
// To generate the specialized mallocgc functions, do 'go run .' inside runtime/_mkmalloc.
//
// To assemble a mallocgc function, the mallocStub function is cloned, and the call to
// inlinedMalloc is replaced with the inlined body of smallScanNoHeaderStub,
// smallNoScanStub or tinyStub, depending on the parameters being specialized.
//
// The size_ (for the tiny case) and elemsize_, sizeclass_, and noscanint_ (for all three cases)
// identifiers are replaced with the value of the parameter in the specialized case.
// The nextFreeFastStub, nextFreeFastTiny, heapSetTypeNoHeaderStub, and writeHeapBitsSmallStub
// functions are also inlined by _mkmalloc.

package runtime
