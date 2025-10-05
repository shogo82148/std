// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package byteorder provides functions for decoding and encoding
// little and big endian integer types from/to byte slices.
package byteorder

func LEUint16(b []byte) uint16

func LEPutUint16(b []byte, v uint16)

func LEAppendUint16(b []byte, v uint16) []byte

func LEUint32(b []byte) uint32

func LEPutUint32(b []byte, v uint32)

func LEAppendUint32(b []byte, v uint32) []byte

func LEUint64(b []byte) uint64

func LEPutUint64(b []byte, v uint64)

func LEAppendUint64(b []byte, v uint64) []byte

func BEUint16(b []byte) uint16

func BEPutUint16(b []byte, v uint16)

func BEAppendUint16(b []byte, v uint16) []byte

func BEUint32(b []byte) uint32

func BEPutUint32(b []byte, v uint32)

func BEAppendUint32(b []byte, v uint32) []byte

func BEUint64(b []byte) uint64

func BEPutUint64(b []byte, v uint64)

func BEAppendUint64(b []byte, v uint64) []byte
