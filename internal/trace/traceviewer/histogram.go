// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package traceviewer

import (
	"github.com/shogo82148/std/html/template"
	"github.com/shogo82148/std/time"
)

// TimeHistogram is an high-dynamic-range histogram for durations.
type TimeHistogram struct {
	Count                int
	Buckets              []int
	MinBucket, MaxBucket int
}

// Add adds a single sample to the histogram.
func (h *TimeHistogram) Add(d time.Duration)

// BucketMin returns the minimum duration value for a provided bucket.
func (h *TimeHistogram) BucketMin(bucket int) time.Duration

// ToHTML renders the histogram as HTML.
func (h *TimeHistogram) ToHTML(urlmaker func(min, max time.Duration) string) template.HTML
