// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Plan 9 directory marshaling. See intro(5).

package syscall

import "github.com/shogo82148/std/errors"

var (
	ErrShortStat = errors.New("stat buffer too short")
	ErrBadStat   = errors.New("malformed stat buffer")
	ErrBadName   = errors.New("bad character in file name")
)

// A Qid represents a 9P server's unique identification for a file.
type Qid struct {
	Path uint64
	Vers uint32
	Type uint8
}

// A Dir contains the metadata for a file.
type Dir struct {
	Type uint16
	Dev  uint32

	Qid    Qid
	Mode   uint32
	Atime  uint32
	Mtime  uint32
	Length int64
	Name   string
	Uid    string
	Gid    string
	Muid   string
}

// Null assigns special "don't touch" values to members of d to
// avoid modifying them during syscall.Wstat.
func (d *Dir) Null()

// Marshal encodes a 9P stat message corresponding to d into b
//
// If there isn't enough space in b for a stat message, ErrShortStat is returned.
func (d *Dir) Marshal(b []byte) (n int, err error)

// UnmarshalDir decodes a single 9P stat message from b and returns the resulting Dir.
//
// If b is too small to hold a valid stat message, ErrShortStat is returned.
//
// If the stat message itself is invalid, ErrBadStat is returned.
func UnmarshalDir(b []byte) (*Dir, error)
