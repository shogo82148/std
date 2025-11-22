// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package byteorder

func LEUint16(b []byte) uint16

func BEUint32(b []byte) uint32

func BEUint64(b []byte) uint64

func LEUint64(b []byte) uint64

func BEPutUint16(b []byte, v uint16)

func BEPutUint32(b []byte, v uint32)

func BEPutUint64(b []byte, v uint64)

func LEPutUint16(b []byte, v uint16)

func LEPutUint64(b []byte, v uint64)

func BEAppendUint16(b []byte, v uint16) []byte

func BEAppendUint32(b []byte, v uint32) []byte

func BEAppendUint64(b []byte, v uint64) []byte
