// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Invariants:
//  At least one of c.sendq and c.recvq is empty.
// For buffered channels, also:
//  c.qcount > 0 implies that c.recvq is empty.
//  c.qcount < c.dataqsiz implies that c.sendq is empty.
