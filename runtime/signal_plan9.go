// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Incoming notes are compared against this table using strncmp, so the
// order matters: longer patterns must appear before their prefixes.
// There are _SIG constants in os2_plan9.go for the table index of some
// of these.
//
// If you add entries to this table, you must respect the prefix ordering
// and also update the constant values is os2_plan9.go.
