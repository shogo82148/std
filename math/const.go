// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージ math は基本的な定数と数学的な関数を提供します。
//
// このパッケージはアーキテクチャ間でビットが完全に同じ結果になることを保証しません。
package math

// 数学定数。
const (
	E   = 2.71828182845904523536028747135266249775724709369995957496696763
	Pi  = 3.14159265358979323846264338327950288419716939937510582097494459
	Phi = 1.61803398874989484820458683436563811772030917980576286213544862

	Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974
	SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931
	SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779
	SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038

	Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009
	Log2E  = 1 / Ln2
	Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790
	Log10E = 1 / Ln10
)

// 浮動小数点の制限値。
// Max はその型で表現可能な最大の有限値です。
// SmallestNonzero はその型で表現可能な最小の正の0以外の値です。
const (
	MaxFloat32             = 0x1p127 * (1 + (1 - 0x1p-23))
	SmallestNonzeroFloat32 = 0x1p-126 * 0x1p-23

	MaxFloat64             = 0x1p1023 * (1 + (1 - 0x1p-52))
	SmallestNonzeroFloat64 = 0x1p-1022 * 0x1p-52
)

// 整数の上限値。
const (
	MaxInt    = 1<<(intSize-1) - 1
	MinInt    = -1 << (intSize - 1)
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint   = 1<<intSize - 1
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)
