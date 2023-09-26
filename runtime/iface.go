// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

//Note: change the formula in the mallocgc call in itabAdd if you change these fields.

// staticbytes is used to avoid convT2E for byte-sized values.
