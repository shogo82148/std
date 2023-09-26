// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// See https://github.com/catapult-project/catapult/blob/master/tracing/docs/embedding-trace-viewer.md
// This is almost verbatim copy of:
// https://github.com/catapult-project/catapult/blob/master/tracing/bin/index.html
// on revision 5f9e4c3eaa555bdef18218a89f38c768303b7b6e.

type Range struct {
	Name  string
	Start int
	End   int
}

type ViewerData struct {
	Events   []*ViewerEvent         `json:"traceEvents"`
	Frames   map[string]ViewerFrame `json:"stackFrames"`
	TimeUnit string                 `json:"displayTimeUnit"`

	footer int
}

type ViewerEvent struct {
	Name     string      `json:"name,omitempty"`
	Phase    string      `json:"ph"`
	Scope    string      `json:"s,omitempty"`
	Time     float64     `json:"ts"`
	Dur      float64     `json:"dur,omitempty"`
	Pid      uint64      `json:"pid"`
	Tid      uint64      `json:"tid"`
	ID       uint64      `json:"id,omitempty"`
	Stack    int         `json:"sf,omitempty"`
	EndStack int         `json:"esf,omitempty"`
	Arg      interface{} `json:"args,omitempty"`
}

type ViewerFrame struct {
	Name   string `json:"name"`
	Parent int    `json:"parent,omitempty"`
}

type NameArg struct {
	Name string `json:"name"`
}

type SortIndexArg struct {
	Index int `json:"sort_index"`
}
