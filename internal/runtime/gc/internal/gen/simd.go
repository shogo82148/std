// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

type Uint8x64 struct {
	valAny
}

func ConstUint8x64(c [64]uint8, name string) (y Uint8x64)

func (x Uint8x64) ToUint64x8() (z Uint64x8)

func (x Uint8x64) Shuffle(shuf Uint8x64) (y Uint8x64)

func (x Uint8x64) ShuffleZeroed(shuf Uint8x64, mask Mask64) (y Uint8x64)

func (x Uint8x64) ShuffleMasked(shuf Uint8x64, mask Mask64) (y Uint8x64)

func (x Uint8x64) Shuffle2(y Uint8x64, shuf Uint8x64) (z Uint8x64)

func (x Uint8x64) Shuffle2Zeroed(y Uint8x64, shuf Uint8x64, mask Mask64) (z Uint8x64)

func (x Uint8x64) Shuffle2Masked(y Uint8x64, shuf Uint8x64, mask Mask64) (z Uint8x64)

type Uint64x8 struct {
	valAny
}

func ConstUint64x8(c [8]uint64, name string) (y Uint64x8)

func BroadcastUint64x8Zeroed(src Uint64, mask Mask8) (z Uint64x8)

func (x Uint64x8) BroadcastMasked(src Uint64, mask Mask8) (z Uint64x8)

func (x Uint64x8) Or(y Uint64x8) (z Uint64x8)

func (x Uint64x8) Sub(y Uint64x8) (z Uint64x8)

func (x Uint64x8) ToUint8x64() (z Uint8x64)

func (x Uint64x8) GF2P8Affine(y Uint8x64) (z Uint8x64)

func (x Uint64x8) ShuffleBits(y Uint8x64) (z Mask64)

func (x Uint64x8) ShuffleBitsMasked(y Uint8x64, mask Mask64) (z Mask64)

type Mask8 struct {
	valAny
}

func ConstMask8(c uint8) (y Mask8)

func (x Mask8) ToUint8() (z Uint64)

func (x Mask8) Or(y Mask8) (z Mask8)

func (x Mask8) ShiftLeft(c uint8) (z Mask8)

type Mask64 struct {
	valAny
}

func ConstMask64(c uint64) (y Mask64)

func (x Mask64) ToUint64() (z Uint64)

func (x Mask64) Or(y Mask64) (z Mask64)

func (x Mask64) ShiftLeft(c uint8) (z Mask64)

func (x Mask64) ShiftRight(c uint8) (z Mask64)
