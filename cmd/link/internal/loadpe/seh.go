// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loadpe

const (
	UNW_FLAG_EHANDLER  = 1 << 3
	UNW_FLAG_UHANDLER  = 2 << 3
	UNW_FLAG_CHAININFO = 3 << 3
)
