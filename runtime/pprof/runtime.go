// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pprof

import (
	"github.com/shogo82148/std/context"
)

// SetGoroutineLabels sets the current goroutine's labels to match ctx.
// A new goroutine inherits the labels of the goroutine that created it.
// This is a lower-level API than [Do], which should be used instead when possible.
func SetGoroutineLabels(ctx context.Context)

// Do calls f with a copy of the parent context with the
// given labels added to the parent's label map.
// Goroutines spawned while executing f will inherit the augmented label-set.
// Each key/value pair in labels is inserted into the label map in the
// order provided, overriding any previous value for the same key.
// The augmented label map will be set for the duration of the call to f
// and restored once f returns.
func Do(ctx context.Context, labels LabelSet, f func(context.Context))
