// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// インターフェースはネットワークインタフェースの名前とインデックスのマッピングを表します。また、ネットワークインタフェースの機能情報も表します。
type Interface struct {
	Index        int
	MTU          int
	Name         string
	HardwareAddr HardwareAddr
	Flags        Flags
}

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
	FlagRunning
)

func (f Flags) String() string

// Addrsは特定のインターフェースのユニキャストインターフェースアドレスのリストを返します。
func (ifi *Interface) Addrs() ([]Addr, error)

// MulticastAddrsは特定のインターフェースに結合されたマルチキャストグループアドレスのリストを返します。
func (ifi *Interface) MulticastAddrs() ([]Addr, error)

// インタフェースはシステムのネットワークインタフェースのリストを返します。
func Interfaces() ([]Interface, error)

// InterfaceAddrsはシステムのユニキャストインターフェースのアドレスのリストを返します。
//
// 返されたリストは関連するインターフェースを識別しません。詳細についてはInterfacesと [Interface.Addrs] を使用してください。
func InterfaceAddrs() ([]Addr, error)

// InterfaceByIndex は、インデックスで指定されたインターフェースを返します。
//
// Solarisでは、論理データリンクを共有する論理ネットワークインターフェースのうちの1つを返しますが、より正確な情報が必要な場合は、
// [InterfaceByName] を使用してください。
func InterfaceByIndex(index int) (*Interface, error)

// InterfaceByNameは、指定された名前で指定されたインターフェースを返します。
func InterfaceByName(name string) (*Interface, error)
