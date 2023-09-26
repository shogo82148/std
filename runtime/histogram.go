// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// timeHistogram represents a distribution of durations in
// nanoseconds.
//
// The accuracy and range of the histogram is defined by the
// timeHistSubBucketBits and timeHistNumBuckets constants.
//
// It is an HDR histogram with exponentially-distributed
// buckets and linearly distributed sub-buckets.
//
// The histogram is safe for concurrent reads and writes.
