// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sql

import (
	"database/sql/driver"
)

var _ driver.DriverContext = &fakeDriverCtx{}

var _ driver.Validator = (*fakeConn)(nil)
