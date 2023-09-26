// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

// testFile is an io.ReadWriteSeeker where the IO operations performed
// on it must match the list of operations in ops.
