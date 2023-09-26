// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sql

import (
	"database/sql/driver"
)

var _ driver.NamedValueChecker = &nvcConn{}

var (
	_ driver.Conn           = &ctxOnlyConn{}
	_ driver.QueryerContext = &ctxOnlyConn{}
	_ driver.ExecerContext  = &ctxOnlyConn{}
)

// badConn implements a bad driver.Conn, for TestBadDriver.
// The Exec method panics.

// badDriver is a driver.Driver that uses badConn.

var _ driver.Pinger = pingConn{}
