// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

const (
	ScorePhi       = iota
	ScoreArg
	ScoreInitMem
	ScoreReadTuple
	ScoreNilCheck
	ScoreMemory
	ScoreReadFlags
	ScoreDefault
	ScoreFlags
	ScoreControl
)

type ValHeap struct {
	a           []*Value
	score       []int8
	inBlockUses []bool
}

func (h ValHeap) Len() int
func (h ValHeap) Swap(i, j int)

func (h *ValHeap) Push(x interface{})

func (h *ValHeap) Pop() interface{}

func (h ValHeap) Less(i, j int) bool
