// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt

// flags placed in a separate struct for easy clearing.

// A fmt is the raw formatter used by Printf etc.
// It prints into a buffer that must be set up separately.
