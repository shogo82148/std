// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pprof

// A profMap is a map from (stack, tag) to mapEntry.
// It grows without bound, but that's assumed to be OK.

// A profMapEntry is a single entry in the profMap.
