// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jpeg

// zigzag maps from the natural ordering to the zig-zag ordering. For example,
// zigzag[0*8 + 3] is the zig-zag sequence number of the element in the fourth
// column and first row.

// unscaledQuantInNaturalOrder are the unscaled quantization tables in
// natural (not zig-zag) order, as specified in section K.1.
