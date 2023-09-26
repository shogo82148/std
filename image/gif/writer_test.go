// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gif

import (
	_ "image/png"
)

// lzw.NewWriter wants an interface which is basically the same thing as gif's
// writer interface.  This ensures we're compatible.
var _ writer = blockWriter{}
