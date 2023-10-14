// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slicewriter

// WriteSeeker is a helper object that implements the io.WriteSeeker
// interface. Clients can create a WriteSeeker, make a series of Write
// calls to add data to it (and possibly Seek calls to update
// previously written portions), then finally invoke BytesWritten() to
// get a pointer to the constructed byte slice.
type WriteSeeker struct {
	payload []byte
	off     int64
}

func (sws *WriteSeeker) Write(p []byte) (n int, err error)

// Seek repositions the read/write position of the WriteSeeker within
// its internally maintained slice. Note that it is not possible to
// expand the size of the slice using SEEK_SET; trying to seek outside
// the slice will result in an error.
func (sws *WriteSeeker) Seek(offset int64, whence int) (int64, error)

// BytesWritten returns the underlying byte slice for the WriteSeeker,
// containing the data written to it via Write/Seek calls.
func (sws *WriteSeeker) BytesWritten() []byte

func (sws *WriteSeeker) Read(p []byte) (n int, err error)
