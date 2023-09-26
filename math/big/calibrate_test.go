// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file prints execution times for the Mul benchmark
// given different Karatsuba thresholds. The result may be
// used to manually fine-tune the threshold constant. The
// results are somewhat fragile; use repeated runs to get
// a clear picture.

// Usage: go test -run=TestCalibrate -calibrate

package big
