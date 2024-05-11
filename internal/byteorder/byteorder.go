// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package byteorder provides functions for decoding and encoding
// little and big endian integer types from/to byte slices.
package byteorder

func LeUint16(b []byte) uint16

func LePutUint16(b []byte, v uint16)

func LeAppendUint16(b []byte, v uint16) []byte

func LeUint32(b []byte) uint32

func LePutUint32(b []byte, v uint32)

func LeAppendUint32(b []byte, v uint32) []byte

func LeUint64(b []byte) uint64

func LePutUint64(b []byte, v uint64)

func LeAppendUint64(b []byte, v uint64) []byte

func BeUint16(b []byte) uint16

func BePutUint16(b []byte, v uint16)

func BeAppendUint16(b []byte, v uint16) []byte

func BeUint32(b []byte) uint32

func BePutUint32(b []byte, v uint32)

func BeAppendUint32(b []byte, v uint32) []byte

func BeUint64(b []byte) uint64

func BePutUint64(b []byte, v uint64)

func BeAppendUint64(b []byte, v uint64) []byte
