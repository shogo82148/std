// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer
// object, creating another object (Reader or Writer) that also implements
// the interface but provides buffering and some help for textual I/O.
package bufio

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)

// Reader implements buffering for an io.Reader object.
type Reader struct {
	buf          []byte
	rd           io.Reader
	r, w         int
	err          error
	lastByte     int
	lastRuneSize int
}

// NewReaderSize returns a new [Reader] whose buffer has at least the specified
// size. If the argument io.Reader is already a [Reader] with large enough
// size, it returns the underlying [Reader].
func NewReaderSize(rd io.Reader, size int) *Reader

// NewReader returns a new [Reader] whose buffer has the default size.
func NewReader(rd io.Reader) *Reader

// Size returns the size of the underlying buffer in bytes.
func (b *Reader) Size() int

// Reset discards any buffered data, resets all state, and switches
// the buffered reader to read from r.
// Calling Reset on the zero value of [Reader] initializes the internal buffer
// to the default size.
// Calling b.Reset(b) (that is, resetting a [Reader] to itself) does nothing.
func (b *Reader) Reset(r io.Reader)

// Peek returns the next n bytes without advancing the reader. The bytes stop
// being valid at the next read call. If Peek returns fewer than n bytes, it
// also returns an error explaining why the read is short. The error is
// [ErrBufferFull] if n is larger than b's buffer size.
//
// Calling Peek prevents a [Reader.UnreadByte] or [Reader.UnreadRune] call from succeeding
// until the next read operation.
func (b *Reader) Peek(n int) ([]byte, error)

// Discard skips the next n bytes, returning the number of bytes discarded.
//
// If Discard skips fewer than n bytes, it also returns an error.
// If 0 <= n <= b.Buffered(), Discard is guaranteed to succeed without
// reading from the underlying io.Reader.
func (b *Reader) Discard(n int) (discarded int, err error)

// Read reads data into p.
// It returns the number of bytes read into p.
// The bytes are taken from at most one Read on the underlying [Reader],
// hence n may be less than len(p).
// To read exactly len(p) bytes, use io.ReadFull(b, p).
// If the underlying [Reader] can return a non-zero count with io.EOF,
// then this Read method can do so as well; see the [io.Reader] docs.
func (b *Reader) Read(p []byte) (n int, err error)

// ReadByte reads and returns a single byte.
// If no byte is available, returns an error.
func (b *Reader) ReadByte() (byte, error)

// UnreadByte unreads the last byte. Only the most recently read byte can be unread.
//
// UnreadByte returns an error if the most recent method called on the
// [Reader] was not a read operation. Notably, [Reader.Peek], [Reader.Discard], and [Reader.WriteTo] are not
// considered read operations.
func (b *Reader) UnreadByte() error

// ReadRune reads a single UTF-8 encoded Unicode character and returns the
// rune and its size in bytes. If the encoded rune is invalid, it consumes one byte
// and returns unicode.ReplacementChar (U+FFFD) with a size of 1.
func (b *Reader) ReadRune() (r rune, size int, err error)

// UnreadRune unreads the last rune. If the most recent method called on
// the [Reader] was not a [Reader.ReadRune], [Reader.UnreadRune] returns an error. (In this
// regard it is stricter than [Reader.UnreadByte], which will unread the last byte
// from any read operation.)
func (b *Reader) UnreadRune() error

// Buffered returns the number of bytes that can be read from the current buffer.
func (b *Reader) Buffered() int

// ReadSlice reads until the first occurrence of delim in the input,
// returning a slice pointing at the bytes in the buffer.
// The bytes stop being valid at the next read.
// If ReadSlice encounters an error before finding a delimiter,
// it returns all the data in the buffer and the error itself (often io.EOF).
// ReadSlice fails with error [ErrBufferFull] if the buffer fills without a delim.
// Because the data returned from ReadSlice will be overwritten
// by the next I/O operation, most clients should use
// [Reader.ReadBytes] or ReadString instead.
// ReadSlice returns err != nil if and only if line does not end in delim.
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

// ReadLine is a low-level line-reading primitive. Most callers should use
// [Reader.ReadBytes]('\n') or [Reader.ReadString]('\n') instead or use a [Scanner].
//
// ReadLine tries to return a single line, not including the end-of-line bytes.
// If the line was too long for the buffer then isPrefix is set and the
// beginning of the line is returned. The rest of the line will be returned
// from future calls. isPrefix will be false when returning the last fragment
// of the line. The returned buffer is only valid until the next call to
// ReadLine. ReadLine either returns a non-nil line or it returns an error,
// never both.
//
// The text returned from ReadLine does not include the line end ("\r\n" or "\n").
// No indication or error is given if the input ends without a final line end.
// Calling [Reader.UnreadByte] after ReadLine will always unread the last byte read
// (possibly a character belonging to the line end) even if that byte is not
// part of the line returned by ReadLine.
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)

// ReadBytes reads until the first occurrence of delim in the input,
// returning a slice containing the data up to and including the delimiter.
// If ReadBytes encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadBytes returns err != nil if and only if the returned data does not end in
// delim.
// For simple uses, a Scanner may be more convenient.
func (b *Reader) ReadBytes(delim byte) ([]byte, error)

// ReadString reads until the first occurrence of delim in the input,
// returning a string containing the data up to and including the delimiter.
// If ReadString encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadString returns err != nil if and only if the returned data does not end in
// delim.
// For simple uses, a Scanner may be more convenient.
func (b *Reader) ReadString(delim byte) (string, error)

// WriteTo implements io.WriterTo.
// This may make multiple calls to the [Reader.Read] method of the underlying [Reader].
// If the underlying reader supports the [Reader.WriteTo] method,
// this calls the underlying [Reader.WriteTo] without buffering.
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)

// Writer implements buffering for an [io.Writer] object.
// If an error occurs writing to a [Writer], no more data will be
// accepted and all subsequent writes, and [Writer.Flush], will return the error.
// After all data has been written, the client should call the
// [Writer.Flush] method to guarantee all data has been forwarded to
// the underlying [io.Writer].
type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}

// NewWriterSize returns a new [Writer] whose buffer has at least the specified
// size. If the argument io.Writer is already a [Writer] with large enough
// size, it returns the underlying [Writer].
func NewWriterSize(w io.Writer, size int) *Writer

// NewWriter returns a new [Writer] whose buffer has the default size.
// If the argument io.Writer is already a [Writer] with large enough buffer size,
// it returns the underlying [Writer].
func NewWriter(w io.Writer) *Writer

// Size returns the size of the underlying buffer in bytes.
func (b *Writer) Size() int

// Reset discards any unflushed buffered data, clears any error, and
// resets b to write its output to w.
// Calling Reset on the zero value of [Writer] initializes the internal buffer
// to the default size.
// Calling w.Reset(w) (that is, resetting a [Writer] to itself) does nothing.
func (b *Writer) Reset(w io.Writer)

// Flush writes any buffered data to the underlying [io.Writer].
func (b *Writer) Flush() error

// Available returns how many bytes are unused in the buffer.
func (b *Writer) Available() int

// AvailableBuffer returns an empty buffer with b.Available() capacity.
// This buffer is intended to be appended to and
// passed to an immediately succeeding [Writer.Write] call.
// The buffer is only valid until the next write operation on b.
func (b *Writer) AvailableBuffer() []byte

// Buffered returns the number of bytes that have been written into the current buffer.
func (b *Writer) Buffered() int

// Write writes the contents of p into the buffer.
// It returns the number of bytes written.
// If nn < len(p), it also returns an error explaining
// why the write is short.
func (b *Writer) Write(p []byte) (nn int, err error)

// WriteByte writes a single byte.
func (b *Writer) WriteByte(c byte) error

// WriteRune writes a single Unicode code point, returning
// the number of bytes written and any error.
func (b *Writer) WriteRune(r rune) (size int, err error)

// WriteString writes a string.
// It returns the number of bytes written.
// If the count is less than len(s), it also returns an error explaining
// why the write is short.
func (b *Writer) WriteString(s string) (int, error)

// ReadFrom implements [io.ReaderFrom]. If the underlying writer
// supports the ReadFrom method, this calls the underlying ReadFrom.
// If there is buffered data and an underlying ReadFrom, this fills
// the buffer and writes it before calling ReadFrom.
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)

// ReadWriter stores pointers to a [Reader] and a [Writer].
// It implements [io.ReadWriter].
type ReadWriter struct {
	*Reader
	*Writer
}

// NewReadWriter allocates a new [ReadWriter] that dispatches to r and w.
func NewReadWriter(r *Reader, w *Writer) *ReadWriter
