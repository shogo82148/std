// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fips

// ResetServiceIndicator clears the service indicator for the running goroutine.
func ResetServiceIndicator()

// ServiceIndicator returns true if and only if all services invoked by this
// goroutine since the last ResetServiceIndicator call are approved.
//
// If ResetServiceIndicator was not called before by this goroutine, its return
// value is undefined.
func ServiceIndicator() bool

// RecordApproved is an internal function that records the use of an approved
// service. It does not override RecordNonApproved calls in the same span.
func RecordApproved()

// RecordNonApproved is an internal function that records the use of a
// non-approved service. It overrides any RecordApproved calls in the same span.
func RecordNonApproved()
