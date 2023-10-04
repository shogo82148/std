// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Range struct {
	Name      string
	Start     int
	End       int
	StartTime int64
	EndTime   int64
}

func (r Range) URL() string

type NameArg struct {
	Name string `json:"name"`
}

type TaskArg struct {
	ID     uint64 `json:"id"`
	StartG uint64 `json:"start_g,omitempty"`
	EndG   uint64 `json:"end_g,omitempty"`
}

type RegionArg struct {
	TaskID uint64 `json:"taskid,omitempty"`
}

type SortIndexArg struct {
	Index int `json:"sort_index"`
}
