// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

// A routingIndex optimizes conflict detection by indexing patterns.
//
// The basic idea is to rule out patterns that cannot conflict with a given
// pattern because they have a different literal in a corresponding segment.
// See the comments in [routingIndex.possiblyConflictingPatterns] for more details.
