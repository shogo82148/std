// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

// testNonEmptyReader wraps an io.Reader and ensures that
// Read is never called with an empty buffer.
