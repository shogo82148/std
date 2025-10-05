// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TestCalibrate determines appropriate thresholds for when to use
// different calculation algorithms. To run it, use:
//
//	go test -run=Calibrate -calibrate >cal.log
//
// Calibration data is printed in CSV format, along with the normal test output.
// See calibrate.md for more details about using the output.

package big
