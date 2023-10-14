// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package browser provides utilities for interacting with users' browsers.
package browser

// Commands returns a list of possible commands to use to open a url.
func Commands() [][]string

// Open tries to open url in a browser and reports whether it succeeded.
func Open(url string) bool
