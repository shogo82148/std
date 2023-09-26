// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP file system request handler

package http

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/time"
)

// A Dir implements http.FileSystem using the native file
// system restricted to a specific directory tree.
//
// An empty Dir is treated as ".".
type Dir string

func (d Dir) Open(name string) (File, error)

// A FileSystem implements access to a collection of named files.
// The elements in a file path are separated by slash ('/', U+002F)
// characters, regardless of host operating system convention.
type FileSystem interface {
	Open(name string) (File, error)
}

// A File is returned by a FileSystem's Open method and can be
// served by the FileServer implementation.
type File interface {
	Close() error
	Stat() (os.FileInfo, error)
	Readdir(count int) ([]os.FileInfo, error)
	Read([]byte) (int, error)
	Seek(offset int64, whence int) (int64, error)
}

// ServeContent replies to the request using the content in the
// provided ReadSeeker.  The main benefit of ServeContent over io.Copy
// is that it handles Range requests properly, sets the MIME type, and
// handles If-Modified-Since requests.
//
// If the response's Content-Type header is not set, ServeContent
// first tries to deduce the type from name's file extension and,
// if that fails, falls back to reading the first block of the content
// and passing it to DetectContentType.
// The name is otherwise unused; in particular it can be empty and is
// never sent in the response.
//
// If modtime is not the zero time, ServeContent includes it in a
// Last-Modified header in the response.  If the request includes an
// If-Modified-Since header, ServeContent uses modtime to decide
// whether the content needs to be sent at all.
//
// The content's Seek method must work: ServeContent uses
// a seek to the end of the content to determine its size.
//
// Note that *os.File implements the io.ReadSeeker interface.
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)

// ServeFile replies to the request with the contents of the named file or directory.
func ServeFile(w ResponseWriter, r *Request, name string)

// FileServer returns a handler that serves HTTP requests
// with the contents of the file system rooted at root.
//
// To use the operating system's file system implementation,
// use http.Dir:
//
//	http.Handle("/", http.FileServer(http.Dir("/tmp")))
func FileServer(root FileSystem) Handler

// httpRange specifies the byte range to be sent to the client.

// countingWriter counts how many bytes have been written to it.
