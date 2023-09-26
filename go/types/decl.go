// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// useCycleMarking enables the new coloring-based cycle marking scheme
// for package-level objects. Set this flag to false to disable this
// code quickly and revert to the existing mechanism (and comment out
// some of the new tests in cycles5.src that will fail again).
// TODO(gri) remove this for Go 1.12

// indir is a sentinel type name that is pushed onto the object path
// to indicate an "indirection" in the dependency from one type name
// to the next. For instance, for "type p *p" the object path contains
// p followed by indir, indicating that there's an indirection *p.
// Indirections are used to break type cycles.

// cutCycle is a sentinel type name that is pushed onto the object path
// to indicate that a cycle doesn't actually exist. This is currently
// needed to break cycles formed via method declarations because they
// are type-checked together with their receiver base types. Once methods
// are type-checked separately (see also TODO in Checker.typeDecl), we
// can get rid of this.
