// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tar implements access to tar archives.
// It aims to cover most of the variations, including those produced
// by GNU and BSD tars.
//
// References:
//
//	http://www.freebsd.org/cgi/man.cgi?query=tar&sektion=5
//	http://www.gnu.org/software/tar/manual/html_node/Standard.html
package tar

import "github.com/shogo82148/std/time"

const (

	// Types
	TypeReg           = '0'
	TypeRegA          = '\x00'
	TypeLink          = '1'
	TypeSymlink       = '2'
	TypeChar          = '3'
	TypeBlock         = '4'
	TypeDir           = '5'
	TypeFifo          = '6'
	TypeCont          = '7'
	TypeXHeader       = 'x'
	TypeXGlobalHeader = 'g'
)

// A Header represents a single header in a tar archive.
// Some fields may not be populated.
type Header struct {
	Name       string
	Mode       int64
	Uid        int
	Gid        int
	Size       int64
	ModTime    time.Time
	Typeflag   byte
	Linkname   string
	Uname      string
	Gname      string
	Devmajor   int64
	Devminor   int64
	AccessTime time.Time
	ChangeTime time.Time
}
