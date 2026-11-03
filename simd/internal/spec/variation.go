// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spec

// Platform identifies a specific platform (generally a GOARCH) for
// platform-defined behavior.
type Platform int8

const (
	PlatformAMD64 Platform = iota
	PlatformARM64
	PlatformWASM

	// PlatformAny permits the behavior of any platform (for example, when
	// testing the portable simd package or emulated execution).
	PlatformAny Platform = -1
)

func (p Platform) String() string

// Variation represents a value of type E that may vary between platforms.
//
// Its zero value is the exact zero value of E.
//
// This type is known to specgen.
type Variation[E EltOrMask] struct {
	canon E
	v     variation[E]
}

// Permits reports whether this variation permits value x on plat.
func (v Variation[E]) Permits(x E, plat Platform) bool

// Canon returns a single canonical value for this variation.
func (v Variation[E]) Canon() E

// String returns this variation as a human-readable string.
func (v Variation[E]) String() string

func (pp perPlatformVar[E]) String() string

func (u ulpsVar[E]) String() string

func (a anyOfVar[E]) String() string

func (undefVar[E]) String() string
