// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

// Default returns the default cache to use.
// It never returns nil.
func Default() Cache

// DefaultDir returns the effective GOCACHE setting.
// It returns "off" if the cache is disabled,
// and reports whether the effective value differs from GOCACHE.
func DefaultDir() (string, bool)
