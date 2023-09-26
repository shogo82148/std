// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// indir is a sentinel type name that is pushed onto the object path
// to indicate an "indirection" in the dependency from one type name
// to the next. For instance, for "type p *p" the object path contains
// p followed by indir, indicating that there's an indirection *p.
// Indirections are used to break type cycles.
