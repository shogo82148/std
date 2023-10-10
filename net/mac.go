// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// HardwareAddrは物理的なハードウェアアドレスを表します。
type HardwareAddr []byte

func (a HardwareAddr) String() string

// ParseMACは、次のいずれかの形式で定義されたIEEE 802 MAC-48、EUI-48、EUI-64、または20オクテットのIP over InfiniBandリンク層アドレスとしてsを解析します：
//
//	00:00:5e:00:53:01
//	02:00:5e:10:00:00:00:01
//	00:00:00:00:fe:80:00:00:00:00:00:00:02:00:5e:10:00:00:00:01
//	00-00-5e-00-53-01
//	02-00-5e-10-00-00-00-01
//	00-00-00-00-fe-80-00-00-00-00-00-00-02-00-5e-10-00-00-00-01
//	0000.5e00.5301
//	0200.5e10.0000.0001
//	0000.0000.fe80.0000.0000.0000.0200.5e10.0000.0001
func ParseMAC(s string) (hw HardwareAddr, err error)
