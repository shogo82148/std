// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// MAC address manipulations

package net

// A HardwareAddr represents a physical hardware address.
type HardwareAddr []byte

func (a HardwareAddr) String() string

// ParseMAC parses s as an IEEE 802 MAC-48, EUI-48, or EUI-64 using one of the
// following formats:
//
//	01:23:45:67:89:ab
//	01:23:45:67:89:ab:cd:ef
//	01-23-45-67-89-ab
//	01-23-45-67-89-ab-cd-ef
//	0123.4567.89ab
//	0123.4567.89ab.cdef
func ParseMAC(s string) (hw HardwareAddr, err error)
