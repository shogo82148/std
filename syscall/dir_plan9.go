// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Plan 9 directory marshaling. See intro(5).

package syscall

import (
	"github.com/shogo82148/std/errors"
)

var (
	ErrShortStat = errors.New("stat buffer too short")
	ErrBadStat   = errors.New("malformed stat buffer")
	ErrBadName   = errors.New("bad character in file name")
)

// Qidは、9Pサーバーにおけるファイルの一意な識別子を表します。
type Qid struct {
	Path uint64
	Vers uint32
	Type uint8
}

// Dirはファイルのメタデータを含みます。
type Dir struct {
	// システムによって変更されるデータ
	Type uint16
	Dev  uint32

	// file data
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

// Nullは、[Wstat]の際に変更されないように、dのメンバーに
// 特別な「変更しない」値を割り当てます。
func (d *Dir) Null()

// Marshalは、dに対応する9Pのstatメッセージをbにエンコードします。
//
// bにstatメッセージのための十分な領域がない場合、[ErrShortStat]が返されます。
func (d *Dir) Marshal(b []byte) (n int, err error)

// UnmarshalDirは、bから単一の9Pのstatメッセージをデコードし、結果のDirを返します。
//
// bが有効なstatメッセージを保持するには小さすぎる場合、[ErrShortStat]が返されます。
//
// statメッセージ自体が無効な場合、[ErrBadStat]が返されます。
func UnmarshalDir(b []byte) (*Dir, error)
