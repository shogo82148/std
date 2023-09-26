// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// A header for a Go map.

// A bucket for a Go map.

// A hash iteration structure.
// If you modify hiter, also change cmd/gc/reflect.c to indicate
// the layout of this structure.
