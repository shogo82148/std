// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

type Dir struct {
	Type uint16
	Dev  uint32

	Qid    Qid
	Mode   uint32
	Atime  uint32
	Mtime  uint32
	Length uint64
	Name   string
	Uid    string
	Gid    string
	Muid   string
}

type Qid struct {
	Path uint64
	Vers uint32
	Type uint8
}

// Null assigns members of d with special "don't care" values indicating
// they should not be written by syscall.Wstat.
func (d *Dir) Null()

// UnmarshalDir reads a 9P Stat message from a 9P protocol message stored in b,
// returning the corresponding Dir struct.
func UnmarshalDir(b []byte) (d *Dir, err error)
