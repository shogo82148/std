// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sql

import (
	"log"
)

var _ = log.Printf

// fakeDriver is a fake database that implements Go's driver.Driver
// interface, just for testing.
//
// It speaks a query language that's semantically similar to but
// syntantically different and simpler than SQL.  The syntax is as
// follows:
//
//   WIPE
//   CREATE|<tablename>|<col>=<type>,<col>=<type>,...
//     where types are: "string", [u]int{8,16,32,64}, "bool"
//   INSERT|<tablename>|col=val,col2=val2,col3=?
//   SELECT|<tablename>|projectcol1,projectcol2|filtercol=?,filtercol2=?
//
// When opening a a fakeDriver's database, it starts empty with no
// tables.  All tables and data are stored in memory only.
