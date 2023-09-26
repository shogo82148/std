// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Address range data structure.
//
// This file contains an implementation of a data structure which
// manages ordered address ranges.

package runtime

// addrRange represents a region of address space.
//
// An addrRange must never span a gap in the address space.

// offAddr represents an address in a contiguous view
// of the address space on systems where the address space is
// segmented. On other systems, it's just a normal address.

// atomicOffAddr is like offAddr, but operations on it are atomic.
// It also contains operations to be able to store marked addresses
// to ensure that they're not overridden until they've been seen.

// addrRanges is a data structure holding a collection of ranges of
// address space.
//
// The ranges are coalesced eagerly to reduce the
// number ranges it holds.
//
// The slice backing store for this field is persistentalloc'd
// and thus there is no way to free it.
//
// addrRanges is not thread-safe.
