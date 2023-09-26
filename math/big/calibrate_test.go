// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Calibration used to determine thresholds for using
// different algorithms.  Ideally, this would be converted
// to go generate to create thresholds.go

// This file prints execution times for the Mul benchmark
// given different Karatsuba thresholds. The result may be
// used to manually fine-tune the threshold constant. The
// results are somewhat fragile; use repeated runs to get
// a clear picture.

// Calculates lower and upper thresholds for when basicSqr
// is faster than standard multiplication.

// Usage: go test -run=TestCalibrate -v -calibrate

package big
