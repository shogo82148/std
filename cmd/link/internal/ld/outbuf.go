// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"github.com/shogo82148/std/cmd/internal/sys"
	"github.com/shogo82148/std/cmd/link/internal/loader"
	"github.com/shogo82148/std/os"
)

// OutBuf is a buffered file writer.
//
// It is similar to the Writer in cmd/internal/bio with a few small differences.
//
// First, it tracks the output architecture and uses it to provide
// endian helpers.
//
// Second, it provides a very cheap offset counter that doesn't require
// any system calls to read the value.
//
// Third, it also mmaps the output file (if available). The intended usage is:
//   - Mmap the output file
//   - Write the content
//   - possibly apply any edits in the output buffer
//   - possibly write more content to the file. These writes take place in a heap
//     backed buffer that will get synced to disk.
//   - Munmap the output file
//
// And finally, it provides a mechanism by which you can multithread the
// writing of output files. This mechanism is accomplished by copying a OutBuf,
// and using it in the thread/goroutine.
//
// Parallel OutBuf is intended to be used like:
//
//	func write(out *OutBuf) {
//	  var wg sync.WaitGroup
//	  for i := 0; i < 10; i++ {
//	    wg.Add(1)
//	    view, err := out.View(start[i])
//	    if err != nil {
//	       // handle output
//	       continue
//	    }
//	    go func(out *OutBuf, i int) {
//	      // do output
//	      wg.Done()
//	    }(view, i)
//	  }
//	  wg.Wait()
//	}
type OutBuf struct {
	arch *sys.Arch
	off  int64

	buf  []byte
	heap []byte

	name   string
	f      *os.File
	encbuf [8]byte
	isView bool
}

func (out *OutBuf) Open(name string) error

func NewOutBuf(arch *sys.Arch) *OutBuf

func (out *OutBuf) View(start uint64) *OutBuf

func (out *OutBuf) Close() error

// ErrorClose closes the output file (if any).
// It is supposed to be called only at exit on error, so it doesn't do
// any clean up or buffer flushing, just closes the file.
func (out *OutBuf) ErrorClose()

// Data returns the whole written OutBuf as a byte slice.
func (out *OutBuf) Data() []byte

func (out *OutBuf) SeekSet(p int64)

func (out *OutBuf) Offset() int64

// Write writes the contents of v to the buffer.
func (out *OutBuf) Write(v []byte) (int, error)

func (out *OutBuf) Write8(v uint8)

// WriteByte is an alias for Write8 to fulfill the io.ByteWriter interface.
func (out *OutBuf) WriteByte(v byte) error

func (out *OutBuf) Write16(v uint16)

func (out *OutBuf) Write32(v uint32)

func (out *OutBuf) Write32b(v uint32)

func (out *OutBuf) Write64(v uint64)

func (out *OutBuf) Write64b(v uint64)

func (out *OutBuf) WriteString(s string)

// WriteStringN writes the first n bytes of s.
// If n is larger than len(s) then it is padded with zero bytes.
func (out *OutBuf) WriteStringN(s string, n int)

// WriteStringPad writes the first n bytes of s.
// If n is larger than len(s) then it is padded with the bytes in pad (repeated as needed).
func (out *OutBuf) WriteStringPad(s string, n int, pad []byte)

// WriteSym writes the content of a Symbol, and returns the output buffer
// that we just wrote, so we can apply further edit to the symbol content.
// For generator symbols, it also sets the symbol's Data to the output
// buffer.
func (out *OutBuf) WriteSym(ldr *loader.Loader, s loader.Sym) []byte
