// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw

// slowestRGBA is a draw.Image like image.RGBA, but it is a different type and
// therefore does not trigger the draw.go fastest code paths.
//
// Unlike slowerRGBA, it does not implement the draw.RGBA64Image interface.

// slowerRGBA is a draw.Image like image.RGBA but it is a different type and
// therefore does not trigger the draw.go fastest code paths.
//
// Unlike slowestRGBA, it still implements the draw.RGBA64Image interface.

// embeddedPaletted is an Image that behaves like an *image.Paletted but whose
// type is not *image.Paletted.
