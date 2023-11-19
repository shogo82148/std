// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.exectracer2

// Simple hash table for tracing. Provides a mapping
// between variable-length data and a unique ID. Subsequent
// puts of the same data will return the same ID.
//
// Uses a region-based allocation scheme and assumes that the
// table doesn't ever grow very big.
//
// This is definitely not a general-purpose hash table! It avoids
// doing any high-level Go operations so it's safe to use even in
// sensitive contexts.

package runtime
