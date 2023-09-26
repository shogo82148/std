// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// call is an in-flight or completed singleflight.Do call

// singleflight represents a class of work and forms a namespace in
// which units of work can be executed with duplicate suppression.

// singleflightResult holds the results of Do, so they can be passed
// on a channel.
