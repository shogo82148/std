// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filepath

// HasPrefix exists for historical compatibility and should not be used.
//
// Deprecated: HasPrefix does not respect path boundaries and
// does not ignore case when required.
func HasPrefix(p, prefix string) bool
