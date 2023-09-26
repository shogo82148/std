// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Time-related runtime and pieces of package time.

package runtime

// Package time knows the layout of this structure.
// If this struct changes, adjust ../time/sleep.go:/runtimeTimer.

// Values for the timer status field.

// maxWhen is the maximum value for timer's when field.

// verifyTimers can be set to true to add debugging checks that the
// timer heaps are valid.
