// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Tests that involve both reading and writing.

package zip

// rleBuffer is a run-length-encoded byte buffer.
// It's an io.Writer (like a bytes.Buffer) and also an io.ReaderAt,
// allowing random-access reads.

// fakeHash32 is a dummy Hash32 that always returns 0.

// suffixSaver is an io.Writer & io.ReaderAt that remembers the last 0
// to 'keep' bytes of data written to it. Call Suffix to get the
// suffix bytes.
