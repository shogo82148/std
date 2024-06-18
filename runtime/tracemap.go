// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple append-only thread-safe hash map for tracing.
// Provides a mapping between variable-length data and a
// unique ID. Subsequent puts of the same data will return
// the same ID. The zero value is ready to use.
//
// Uses a region-based allocation scheme internally, and
// reset clears the whole map.
//
// It avoids doing any high-level Go operations so it's safe
// to use even in sensitive contexts.

package runtime
