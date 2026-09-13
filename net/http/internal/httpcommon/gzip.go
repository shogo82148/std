// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpcommon

import (
	"github.com/shogo82148/std/compress/gzip"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
)

// GzipReader wraps a response body so it can lazily
// get gzip.Reader from the pool on the first call to Read.
// After Close is called it puts gzip.Reader to the pool immediately
// if there is no Read in progress or later when Read completes.
type GzipReader struct {
	_    incomparable
	Body io.ReadCloser
	mu   sync.Mutex
	zr   *gzip.Reader
	zerr error
}

func (gz *GzipReader) Read(p []byte) (n int, err error)

func (gz *GzipReader) Close() error
