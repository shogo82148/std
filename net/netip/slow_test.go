// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package netip_test

import (
	. "net/netip"
)

// zeros is a slice of eight stringified zeros. It's used in
// parseIPSlow to construct slices of specific amounts of zero fields,
// from 1 to 8.
