// "go test -run=Generate -write=all"によって生成されたコードです；編集しないでください。

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// Chanはチャネルの型を表します。
type Chan struct {
	dir  ChanDir
	elem Type
}

// ChanDirの値は、チャネルの方向を示します。
type ChanDir int

// チャネルの方向は、次の定数のいずれかで示されます。
const (
	SendRecv ChanDir = iota
	SendOnly
	RecvOnly
)

// NewChanは指定された方向と要素の型のための新しいチャネル型を返します。
func NewChan(dir ChanDir, elem Type) *Chan

// Dir 関数は、チャネル c の方向を返します。
func (c *Chan) Dir() ChanDir

// Elem はチャネル c の要素の型を返します。
func (c *Chan) Elem() Type

func (c *Chan) Underlying() Type
func (c *Chan) String() string
