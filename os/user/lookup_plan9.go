// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package user

// Partial os/user support on Plan 9.
// Supports Current(), but not Lookup()/LookupId().
// The latter two would require parsing /adm/users.
