// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elf

// errorReader returns error from all operations.

// readSeekerFromReader converts an io.Reader into an io.ReadSeeker.
// In general Seek may not be efficient, but it is optimized for
// common cases such as seeking to the end to find the length of the
// data.
