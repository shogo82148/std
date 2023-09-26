// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

// failOnceWriter fails exactly once and then always reports success.

// testNonEmptyWriter wraps an io.Writer and ensures that
// Write is never called with an empty buffer.
