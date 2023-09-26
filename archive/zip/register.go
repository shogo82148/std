// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"github.com/shogo82148/std/io"
)

// A Compressor returns a compressing writer, writing to the
// provided writer. On Close, any pending data should be flushed.
type Compressor func(io.Writer) (io.WriteCloser, error)

// Decompressor is a function that wraps a Reader with a decompressing Reader.
// The decompressed ReadCloser is returned to callers who open files from
// within the archive.  These callers are responsible for closing this reader
// when they're finished reading.
type Decompressor func(io.Reader) io.ReadCloser

// RegisterDecompressor allows custom decompressors for a specified method ID.
func RegisterDecompressor(method uint16, d Decompressor)

// RegisterCompressor registers custom compressors for a specified method ID.
// The common methods Store and Deflate are built in.
func RegisterCompressor(method uint16, comp Compressor)
