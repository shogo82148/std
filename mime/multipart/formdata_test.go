// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package multipart

// failOnReadAfterErrorReader is an io.Reader wrapping r.
// It fails t if any Read is called after a failing Read.
