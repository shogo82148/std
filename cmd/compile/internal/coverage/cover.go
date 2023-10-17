// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package coverage

// Fixup adds calls to the pkg init function as appropriate to
// register coverage-related variables with the runtime.
//
// It also reclassifies selected variables (for example, tagging
// coverage counter variables with flags so that they can be handled
// properly downstream).
func Fixup()
