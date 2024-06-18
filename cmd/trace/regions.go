// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/net/http"
)

// UserRegionsHandlerFunc returns a HandlerFunc that reports all regions found in the trace.
func UserRegionsHandlerFunc(t *parsedTrace) http.HandlerFunc

// UserRegionHandlerFunc returns a HandlerFunc that presents the details of the selected regions.
func UserRegionHandlerFunc(t *parsedTrace) http.HandlerFunc
