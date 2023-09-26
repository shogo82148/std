// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sql

import (
	"database/sql/driver"
)

// fakeDriver is a fake database that implements Go's driver.Driver
// interface, just for testing.
//
// It speaks a query language that's semantically similar to but
// syntactically different and simpler than SQL.  The syntax is as
// follows:
//
//   WIPE
//   CREATE|<tablename>|<col>=<type>,<col>=<type>,...
//     where types are: "string", [u]int{8,16,32,64}, "bool"
//   INSERT|<tablename>|col=val,col2=val2,col3=?
//   SELECT|<tablename>|projectcol1,projectcol2|filtercol=?,filtercol2=?
//   SELECT|<tablename>|projectcol1,projectcol2|filtercol=?param1,filtercol2=?param2
//
// Any of these can be preceded by PANIC|<method>|, to cause the
// named method on fakeStmt to panic.
//
// Any of these can be proceeded by WAIT|<duration>|, to cause the
// named method on fakeStmt to sleep for the specified duration.
//
// Multiple of these can be combined when separated with a semicolon.
//
// When opening a fakeDriver's database, it starts empty with no
// tables. All tables and data are stored in memory only.

var _ driver.DriverContext = &fakeDriverCtx{}

type Dummy struct {
	driver.Driver
}

// hook to simulate connection failures

var _ driver.Validator = (*fakeConn)(nil)

// hook to simulate broken connections

// hook to simulate broken connections

// hook to simulate broken connections

// hook to simulate broken connections

// hook to simulate broken connections

// fakeDriverString is like driver.String, but indirects pointers like
// DefaultValueConverter.
//
// This could be surprising behavior to retroactively apply to
// driver.String now that Go1 is out, but this is convenient for
// our TestPointerParamsAndScans.
//
