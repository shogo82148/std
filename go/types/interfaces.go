// Copyright 2017 The Go Authors. All rights reserved.
// Use of this src code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// ifaceInfo describes the method set for an interface.
// The zero value for an ifaceInfo is a ready-to-use ifaceInfo representing
// the empty interface.

// emptyIfaceInfo represents the ifaceInfo for the empty interface.

// methodInfo represents an interface method.
// At least one of src or fun must be non-nil.
// (Methods declared in the current package have a non-nil scope
// and src, and eventually a non-nil fun field; imported and pre-
// declared methods have a nil scope and src, and only a non-nil
// fun field.)

// A methodInfoSet maps method ids to methodInfos.
// It is used to determine duplicate declarations.
// (A methodInfo set is the equivalent of an objset
// but for methodInfos rather than Objects.)
