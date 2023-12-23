// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Serving of pprof-like profiles.

package traceviewer

import (
	"github.com/shogo82148/std/internal/profile"
	"github.com/shogo82148/std/internal/trace"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/time"
)

type ProfileFunc func(r *http.Request) ([]ProfileRecord, error)

// SVGProfileHandlerFunc serves pprof-like profile generated by prof as svg.
func SVGProfileHandlerFunc(f ProfileFunc) http.HandlerFunc

type ProfileRecord struct {
	Stack []*trace.Frame
	Count uint64
	Time  time.Duration
}

func BuildProfile(prof []ProfileRecord) *profile.Profile